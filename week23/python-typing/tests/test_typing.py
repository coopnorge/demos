# import logging
import random
from pathlib import Path
from typing import (
    Any,
    Dict,
    Generic,
    List,
    Optional,
    Protocol,
    Set,
    Tuple,
    TypeVar,
    Union,
    runtime_checkable,
)

import pytest
from attr import dataclass


def test_hints() -> None:
    some_int = 3
    with pytest.raises(TypeError):
        some_int += "13"

    class Foo:
        def __init__(self, input) -> None:
            self.value = "one" + input

    with pytest.raises(TypeError):
        some_foo = Foo(3)

    some_foo = Foo("3")
    with pytest.raises(AttributeError):
        some_foo.valye += "14"


def test_scalars() -> None:
    my_int: int = 3
    my_int = 3
    my_int = True
    assert isinstance(my_int, int)
    my_int = "str"
    my_int = 1.1

    my_bool = True
    my_bool = False
    my_bool = "str"
    my_str = "str"
    my_str = "other"
    my_float = 1.1
    my_float = 1.2


def test_optional() -> None:
    my_value: Optional[int]

    # always 3
    my_value = 3
    my_value += 6

    # always None
    my_value = None
    assert not isinstance(my_value, int)

    # always None
    my_value = 3 if random.choice([False, False]) else None
    with pytest.raises(TypeError):
        my_value += 6

    # always 3
    my_value = 3 if random.choice([True, True]) else None
    assert my_value is not None
    my_value += 6

    # sometimes 3, sometimes None
    my_value = 3 if random.choice([False, True]) else None
    if my_value is not None:
        my_value += 6


def test_func() -> None:
    def my_func(a: int, b: Optional[str] = None) -> str:
        bv = b if b is not None else "###"
        return f"{a + 3}{'.' + bv}"

    assert my_func(3) == "6.###"
    assert my_func(3, "six") == "6.six"

    with pytest.raises(TypeError):
        my_func("wrong")

    with pytest.raises(TypeError):
        my_func(z=33)

    with pytest.raises(TypeError):
        my_func(3, 7)

    with pytest.raises(TypeError):
        my_func(3, 7)


@dataclass(frozen=True)
class Point:
    x: float
    y: float


def test_dataclass() -> None:
    with pytest.raises(TypeError):
        point = Point(a=1, b=2)

    point = Point("1", "2")


def test_any() -> None:
    value: Any = random.choice((1, 1.111, Point(9, 9), "one"))

    with pytest.raises(AttributeError):
        value.missing(1)

    if isinstance(value, float):
        float_val = value + 3
        float_val = "3"
    elif isinstance(value, int):
        int_val = value + 3
        int_val = "3"
    elif isinstance(value, Point):
        point_val = value.x + value.y
        point_val = "3"
    elif isinstance(value, str):
        str_val = value + "string"
        str_val = 1


def test_union() -> None:
    ValueT = Union[int, float, Point, str]
    values: List[ValueT] = [3, 3.33, Point(4, 4), "111", Path("/dev/")]

    for value in values:
        if isinstance(value, float):
            float_val = value + 3
            float_val = "3"
        elif isinstance(value, int):
            int_val = value + 3
            int_val = "3"
        elif isinstance(value, Point):
            point_val = value.x + value.y
            point_val = "3"
        elif isinstance(value, str):
            str_val = value + "string"
            str_val = 1
        elif isinstance(value, Path):
            path_val = value / "null"
            path_val = 1


def test_set() -> None:
    # set
    my_set: Set[Point] = {Point(1, 1), Point(2, 2)}
    my_set.add(3)
    my_set.add("nine")


def test_tuple() -> None:
    my_tuple: Tuple[str, int, float] = ("one", 2, 3.33)
    my_tuple = (3.33, "one", 2)


def test_dict() -> None:
    # dict
    my_dict: Dict[str, int] = {
        "3": 3,
        "4": 4,
    }
    my_dict[9] = "a"


def test_generic_func() -> None:
    AnyTypeVar = TypeVar("AnyTypeVar")

    def coalesce(*args: Optional[AnyTypeVar]) -> Optional[AnyTypeVar]:
        for arg in args:
            if arg is not None:
                return arg
        return None

    coalesced = coalesce(1, None, None, 3, "str")


def test_generic_func_bound() -> None:
    IntTypeVar = TypeVar("IntTypeVar", bound=int)

    def coalesce(*args: Optional[IntTypeVar]) -> Optional[IntTypeVar]:
        for arg in args:
            if arg is not None:
                return arg
        return None

    coalesced = coalesce(1, None, None, False, "str")


def test_generic_class() -> None:
    AnyTypeVar = TypeVar("AnyTypeVar")

    @dataclass(frozen=True)
    class Holder(Generic[AnyTypeVar]):
        value: AnyTypeVar

        def get_value(self) -> AnyTypeVar:
            return self.value

        def defined(self) -> bool:
            return self.value is None

    h_int = Holder(1)
    z = h_int.get_value() + 1
    with pytest.raises(TypeError):
        z = h_int.get_value() + "1"


@runtime_checkable
class Greeter(Protocol):
    @property
    def name(self) -> str:
        ...

    def greet(self) -> str:
        ...


@dataclass
class HelloWorld:
    name = "opener"

    def greet(self) -> str:
        return f"{self.name} says hello"


@dataclass
class Dog:
    name = "kaptein"

    def bark(self) -> str:
        return f"{self.name} barks"


@dataclass
class Parrot(Greeter):
    name = "polly"

    def greet(self, something: Any) -> None:
        return None


def test_protocol() -> None:
    def use_greeter(greeter: Greeter) -> str:
        return greeter.greet()

    assert use_greeter(HelloWorld()) == f"{HelloWorld.name} says hello"
    with pytest.raises(AttributeError):
        assert use_greeter(Dog()) == f"{HelloWorld.name} says hello"


# Final notes:
# - https://docs.python.org/3/library/typing.html#typing.NoReturn
# - https://docs.python.org/3/library/typing.html#newtype
# - https://docs.python.org/3/library/typing.html#callable
# - https://docs.python.org/3/library/typing.html#typing.Literal
# - https://docs.python.org/3/library/typing.html#typing.TypedDict
