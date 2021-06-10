import json
import logging
from typing import Type

import jsonschema  # type: ignore[import]
import jsonschema.exceptions  # type: ignore[import]
import pytest
import yaml
from pydantic import BaseModel, Field
from pydantic.types import conint
from pytest_subtests import SubTests  # type: ignore[import]
from typing_extensions import Annotated


class Point(BaseModel):
    x: float
    y: float


class PointQ1(Point):
    x: Annotated[float, Field(gt=0)]
    y: Annotated[float, Field(gt=0)]


def test_point(subtests: SubTests) -> None:
    p: Point = Point(x=-1, y=-1)
    with pytest.raises(ValueError):
        p = Point(a=1, b=1)

    with pytest.raises(ValueError):
        p = PointQ1(x=-1, y=-1)

    xy = p.x + p.y

    p = PointQ1(x=1, y=1)
    assert (
        json.dumps(PointQ1.schema(), indent="  ", sort_keys=True)
        == """\
{
  "properties": {
    "x": {
      "exclusiveMinimum": 0,
      "title": "X",
      "type": "number"
    },
    "y": {
      "exclusiveMinimum": 0,
      "title": "Y",
      "type": "number"
    }
  },
  "required": [
    "x",
    "y"
  ],
  "title": "PointQ1",
  "type": "object"
}"""
    )

    p_yaml = "{ x: -1, y: 1 }"

    jsonschema.validate(yaml.safe_load(p_yaml), schema=Point.schema())

    with pytest.raises(jsonschema.exceptions.ValidationError):
        jsonschema.validate(yaml.safe_load(p_yaml), schema=PointQ1.schema())

    p = Point.parse_obj(yaml.safe_load(p_yaml))
    n = p.x
    n = p.y
    with pytest.raises(AttributeError):
        n = p.z
