# NMM in OWL Demo

From [OWL 2 Web Ontology Language Primer](https://www.w3.org/TR/owl2-primer/):

## What

> An ontology is a set of precise descriptive statements about some part of the world (usually referred to as the domain of interest or the subject matter of the ontology).

Simpler:
* What types/classes exist
* What are the relationships beteen these types
* What properties/atributes can the types have
* What are the constraints on types and properties.
* Machine readable, queryable, meta-circular/self-hosting.

## Why

* Formal semantics for modelling data/knowledge
* Prevent misunderstanding between humans and machines.
* Accomodate an open world with no unique name assumption
* Accomodate heterogeneity
  * Format and System heterogeneity: not tied to one data format or data source. Accomodates JSON, CSV, Relational DBs, Document DBs, ...
  * Schematic and Structural heterogeneity: not tied to one schema for the data, different scheas can be comined effectively.
  * Semantic heterogeneity: Different ontologies that assign different semantics to the same types and properties


## Used vocabularies/ontologies

* org: [The Organization Ontology (W3C)](https://www.w3.org/TR/vocab-org/)
* foaf: [FOAF Vocabulary Specification 0.99](http://xmlns.com/foaf/spec/)
* rdau: [RDA element sets unconstrained (RDA)](http://www.rdaregistry.info/Elements/u/)
** ownership relations

## See Also

* [RDF Schema 1.1](https://www.w3.org/TR/rdf-schema/)
* https://en.wikipedia.org/wiki/Open-world_assumption
* https://en.wikipedia.org/wiki/Unique_name_assumption
* https://www.w3.org/TR/csv2rdf/#example-events-listing
* https://a.ml/
* https://solidproject.org/

## What it is not

* Not a document/graph validation language (see [SHACL](https://www.w3.org/TR/shacl/))
* Data format

## Using ...

```bash
make toolchain
make
live-server --port=27480 --open=webvowl/#coop-ontology.vowl build/
```

- http://localhost:27480/webvowl/#coop-ontology.vowl
- http://localhost:27480/webvowl/#dataprotection-rdf.vowl
- http://localhost:27480/webvowl/#acl.vowl
- http://localhost:27480/coop-ontology.owl.pylode.html 
- http://localhost:27480/coop-ontology.owl.widoco/doc/index-en.html

## Resources

* [Linked Open Vocabularies (LOV)](https://lov.linkeddata.es/dataset/lov/)
  Finding vocabularies
* [zazuko prefix server](https://prefix.zazuko.com/)
  Resolve RDF Terms

## Queries

```bash
## All types
poetry run rdflib-xtl sparql --if turtle \
  -q 'PREFIX coopo:  <xri://coop.no/tbd-20210426/ontology/> 
    SELECT ?s ?label WHERE {
        ?s rdfs:subClassOf owl:Thing.
        ?s rdfs:isDefinedBy coopo:.
        ?s rdfs:label ?label.
    }' \
  build/coop-ontology.owl.rsowlrl.ttl | column -s, -t

## All subclasses of people

poetry run rdflib-xtl sparql --if turtle \
  -q 'PREFIX coopo:  <xri://coop.no/tbd-20210426/ontology/> 
    SELECT ?s ?label WHERE {
        ?s rdfs:subClassOf foaf:Person.
        ?s rdfs:isDefinedBy coopo:.
        ?s rdfs:label ?label.
    }' \
  build/coop-ontology.owl.rsowlrl.ttl | column -s, -t

## All properties of people

poetry run rdflib-xtl sparql --if turtle \
  -q 'PREFIX coopo:  <xri://coop.no/tbd-20210426/ontology/> 
    SELECT ?s ?label WHERE {
        ?s rdf:type rdf:Property.
        ?s rdfs:domain foaf:Person.
        ?s rdfs:label ?label.
    }' \
  build/coop-ontology.owl.rsowlrl.ttl | column -s, -t

## All properties of groups

poetry run rdflib-xtl sparql --if turtle \
  -q 'PREFIX coopo:  <xri://coop.no/tbd-20210426/ontology/> 
    SELECT ?s ?label WHERE {
        ?s rdf:type rdf:Property.
        ?s rdfs:domain foaf:Group.
        ?s rdfs:label ?label.
    }' \
  build/coop-ontology.owl.rsowlrl.ttl | column -s, -t

## All identifiers of entities

poetry run rdflib-xtl sparql --if turtle \
  -q 'PREFIX coopo:  <xri://coop.no/tbd-20210426/ontology/> 
    SELECT ?s ?label WHERE {
        ?s rdf:type rdf:Property.
        ?s rdfs:subPropertyOf dc:identifier.
        ?s rdfs:label ?label.
    }' \
  build/coop-ontology.owl.rsowlrl.ttl | column -s, -t

## All identifiers of people

poetry run rdflib-xtl sparql --if turtle \
  -q 'PREFIX coopo:  <xri://coop.no/tbd-20210426/ontology/> 
    SELECT ?s ?label WHERE {
        ?s rdf:type rdf:Property.
        ?s rdfs:subPropertyOf dc:identifier.
        ?s rdfs:domain foaf:Person.
        ?s rdfs:label ?label.
    }' \
  build/coop-ontology.owl.rsowlrl.ttl | column -s, -t
```
