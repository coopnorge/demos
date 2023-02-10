# vim: set noexpandtab fo-=t:

# https://www.gnu.org/software/make/manual/make.html

########################################################################
# boiler plate
########################################################################

.DELETE_ON_ERROR:

.PHONY: default
default:

SHELL=bash
current_makefile:=$(lastword $(MAKEFILE_LIST))
current_makefile_dirname:=$(dir $(current_makefile))
current_makefile_dirname_abspath:=$(dir $(abspath $(current_makefile)))
current_makefile_dirname_realpath:=$(dir $(realpath $(current_makefile)))

__empty:=
__space:=$(__empty) $(__empty)

define __newline


endef

ifneq ($(filter all vars,$(VERBOSE)),)
dump_var=$(info var $(1)=$($(1)))
dump_vars=$(foreach var,$(1),$(call dump_var,$(var)))
else
dump_var=
dump_vars=
endif

ifneq ($(filter all targets,$(VERBOSE)),)
__ORIGINAL_SHELL:=$(SHELL)
SHELL=$(warning Building $@$(if $<, (from $<))$(if $?, ($? newer)))$(TIME) $(__ORIGINAL_SHELL)
endif

$(call dump_vars,current_makefile current_makefile_dirname \
	current_makefile_dirname_abspath current_makefile_dirname_realpath)

########################################################################
# useful ...
########################################################################

## force ...
.PHONY: .FORCE
.FORCE:
$(force_targets): .FORCE

## dirs ...
.PRECIOUS: %/
%/:
	mkdir -vp $(@)

.PHONY: clean-%/
clean-%/:
	@{ test -d $(*) && { set -x; rm -vr $(*); set +x; } } || echo "directory $(*) does not exist ... nothing to clean"

## stamps
.PHONY: .FORCE
__refresh_stamps_target=$(if $(refresh_stamps),.FORCE,)

%.stamp: $(__refresh_stamps_target)
	echo stamp > $(@)

%.tstamp: $(__refresh_stamps_target)
	date +%Y-%m-%dT%H:%M:%S > $(@)

.PHONY: clean-stamp clean
clean: clean-stamp
clean-stamp:
	rm -v *.stamp *.tsstamp || true

## functions ...
reflex_path=$(subst $(__space),/,$(foreach dir,$(subst /,$(__space),$(1)),..))

true:=T
false:=
seq = $(if $(subst x$(1)x,,x$(2)x)$(subst x$(2)x,,x$(1)x),$(false),$(true))
assert=$(if $(call seq,$(1),$(2)),$(1),$(error expected output $(2) did not match actual output $(1)))

########################################################################
# local ...
########################################################################

########################################################################
# general ...
########################################################################

.PHONY: nothing
nothing:
	@echo doing $(@) ...


.PHONY: all
all: default
	@echo doing $(@) ...

.PHONY: clean-build
clean-build:
	rm -vr $(build_dir)/ || true

.PHONY: clean-tools
clean-tools:
	rm -vr $(tools_dir)/ || true

clean: clean-build

build_dir=build
tools_dir=local-tools
cache_dir=$(HOME)/.cache

########################################################################
# sources
########################################################################

yaml_sources=$(filter-out $(build_dir)/% trash/%,$(wildcard *.yaml) $(wildcard */*.yaml) $(wildcard */*/*.yaml))

rdflint_jar=$(tools_dir)/rdflint.jar
rdflint=java -jar $(rdflint_jar)
robot_jar=$(tools_dir)/robot.jar
robot=java -jar $(robot_jar)

jena_dir=$(tools_dir)/apache-jena
jena_riot=$(jena_dir)/bin/riot
jena_schemagen=$(jena_dir)/bin/schemagen

owl2vowl_jar=$(tools_dir)/owl2vowl.jar

widoco_jar=$(tools_dir)/widoco.jar

ontoviz=$(tools_dir)/ontology-visualization/ontology_viz.py


########################################################################
# targets
########################################################################

json_targets=$(foreach source,$(yaml_sources),$(build_dir)/$(source:.yaml=.json))
$(call dump_vars,yaml_sources json_targets)
.PHONY: json
json: $(json_targets)

default: $(build_dir)/coop-ontology.owl.nt
default: $(build_dir)/coop-ontology.owl.ttl
default: $(build_dir)/coop-ontology.owl.xml
default: $(build_dir)/coop-ontology.owl.ld.json
default: $(build_dir)/coop-ontology.owl.unsatisfiable.ttl
default: $(build_dir)/coop-ontology.owl.report.txt
default: $(build_dir)/coop-ontology.owl.rselk.ttl
default: $(build_dir)/coop-ontology.owl.rsowlrl.ttl
default: $(build_dir)/coop-ontology.vowl.json

default: $(build_dir)/coop-ontology.owl.widoco/doc/index-en.html
default: $(build_dir)/coop-ontology.owl.pylode.html

default: $(build_dir)/dataprotection-rdf.vowl.json
default: $(build_dir)/acl.vowl.json
default: $(build_dir)/foaf.vowl.json
default: $(build_dir)/org.vowl.json



default: $(build_dir)/coop-example-a.ld.json
default: $(build_dir)/coop-example-a.ld.yaml

default: $(build_dir)/coop-example-b.ttl
default: $(build_dir)/coop-example-b.nt
default: $(build_dir)/coop-example-b.nq
default: $(build_dir)/coop-example-b.ld.yaml

default: $(build_dir)/coop-ontology-x.owl.ttl
$(build_dir)/coop-ontology-x.owl.ttl: $(build_dir)/coop-ontology.owl.ttl ext/foaf.owl.ttl  ext/org.owl.ttl  ext/rdau.owl.ttl
	$(venv_run) rdfpipe -i turtle -o turtle $(^) > $(@)



########################################################################
# patterns
########################################################################

#venv_run=pipenv run
venv_run=poetry run

$(build_dir)/%.json: %.yaml | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	cat $(<) | $(venv_run) python -c 'import sys, yaml, json; json.dump(yaml.safe_load(sys.stdin), sys.stdout, indent=2)' > $(@)

$(build_dir)/%.ttl: %.ld.yaml $(jena_riot) | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	#cat $(<) | python -c 'import sys, yaml, json; json.dump(yaml.safe_load(sys.stdin), sys.stdout, indent=2)' | $(jena_riot) --syntax JSON-LD --out TURTLE - > $(@)
	cat $(<) | $(venv_run) python -c 'import sys, yaml, json; json.dump(yaml.safe_load(sys.stdin), sys.stdout, indent=2)' | $(venv_run) rdfpipe -i json-ld -o turtle /dev/stdin > $(@)

$(build_dir)/%.ttl: %.ld.json $(jena_riot) | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -i json-ld -o turtle $(<) > $(@)

#.PRECIOUS: $(build_dir)/%.ttl

$(build_dir)/%.ld.json: $(build_dir)/%.ttl $(jena_riot) | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -o json-ld -i turtle $(<) > $(@)

$(build_dir)/%.nq: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -i turtle -o nquads $(<) > $(@)

$(build_dir)/%.nt: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -i turtle -o application/n-triples $(<) > $(@)

$(build_dir)/%.ttl: ext/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -i turtle -o turtle $(<) > $(@)

$(build_dir)/%.ttl: %.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -i turtle -o turtle $(<) > $(@)

$(build_dir)/%.rvz.png: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | neato -Tpng -Gdpi=300 -Goverlap=scale -Esep=1.5 -o $(@) /dev/stdin

$(build_dir)/%.rvz.svg: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | neato -Tsvg -Gdpi=300 -Goverlap=scale -o $(@) /dev/stdin

$(build_dir)/%.rvs.png: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | dot -Tpng -Gdpi=300 -Goverlap=scale $(dot_args) -o $(@) /dev/stdin

$(build_dir)/%.rvs.svg: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | dot -Tsvg -Gdpi=300 -Goverlap=scale $(dot_args) -o $(@) /dev/stdin

$(build_dir)/%.rvs.xdot: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | gawk '(1); /digraph/{ print "  graph[overlap=scale]" };' | $(venv_run) xdot -f dot /dev/stdin

# https://dreampuf.github.io/GraphvizOnline/
# http://magjac.com/graphviz-visual-editor/
# https://graphviz.org/doc/info/attrs.html
# http://viz-js.com/
$(build_dir)/%.rvs.dot: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | gawk '(1); /digraph/{ print "  graph[dpi=300]" };' >$(@)

$(build_dir)/%.rvz.svg: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | neato -Tsvg -Gdpi=300 -Goverlap=scale -o $(@) /dev/stdin

$(build_dir)/%.rvz.xdot: $(build_dir)/%.ttl | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) python local-tools/ontology-visualization/ontology_viz.py $(<) -o /dev/stdout | gawk '(1); /digraph/{ print "  graph[overlap=scale]" };' | $(venv_run) xdot -f neato /dev/stdin

$(build_dir)/%.xml: $(build_dir)/%.ttl $(jena_riot) | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdfpipe -i turtle -o application/rdf+xml '$(<)' >/dev/null
	$(jena_riot) --out RDF/XML '$(<)' > '$(@)'
.PRECIOUS: $(build_dir)/%.xml


$(build_dir)/%.java.done: $(build_dir)/%.owl.xml $(build_dir)/%.sgen.xml | $(jena_schemagen) $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(jena_schemagen) -i $(<) -c $(word 2,$(^))
	touch $(@)

$(build_dir)/%.ld.json: %.ttl $(build_dir)/context.ld.json | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	#riot --out RDF/JSON $(<) | jsonld compact -c $(abspath $(build_dir))/context.ld.json > $(@)
	$(venv_run) rdfpipe -i turtle -o json-ld '$(<)' | jsonld compact -c $(abspath $(build_dir))/context.ld.json > '$(@)'
	#touch $(@)

$(build_dir)/%.ld.yaml: $(build_dir)/%.ld.json
	mkdir -vp $$(dirname -- "$(@)")
	cat '$(<)' | $(venv_run) python -c 'import sys, yaml, json; yaml.safe_dump(json.load(sys.stdin), sys.stdout, indent=2)' > '$(@)'
	#touch $(@)

$(build_dir)/%.owl.widoco/doc/index-en.html: $(build_dir)/%.owl.xml $(widoco_jar)
	java -jar $(widoco_jar) \
		-rewriteAll -licensius -getOntologyMetadata -uniteSections \
		-includeAnnotationProperties \
		-webVowl \
		-ontFile $(<) -outFolder $$(dirname $$(dirname -- "$(@)"))
	@echo google-chrome $(@)

$(build_dir)/%.owl.unsatisfiable.ttl: $(build_dir)/%.owl.ttl | $(robot_jar) $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(robot) reason --reasoner ELK --input '$(<)' --dump-unsatisfiable '$(@)'
	touch '$(@)'

$(build_dir)/%.owl.report.txt: $(build_dir)/%.owl.ttl | $(robot_jar) $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(robot) report --input '$(<)' --output /dev/stdout | grep -v 'missing_definition' | tee '$(@)' 1>&2

$(build_dir)/%.owl.rselk.ttl: $(build_dir)/%.owl.ttl | $(robot_jar) $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(robot) reason --reasoner ELK --axiom-generators "SubClass" --input '$(<)' --output '$(@)'

$(build_dir)/%.owl.rsowlrl.ttl: $(build_dir)/%.owl.ttl | $(robot_jar) $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	$(venv_run) rdflib-xtl reason --if turtle --of turtle '$(<)' > '$(@)'

$(build_dir)/%.vowl.json: $(build_dir)/%.owl.ttl $(owl2vowl_jar) $(build_dir)/webvowl/index.html | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	java -jar $(owl2vowl_jar) -file '$(<)' -output '$(@)'
	cp -vp '$(@)' $(build_dir)/webvowl/data/
	@echo live-server build/ --open='webvowl/#$(*).vowl'

$(build_dir)/%.vowl.json.view: $(build_dir)/%.vowl.json
	live-server build/ --open='webvowl/#$(*).vowl'

$(build_dir)/%.owl.pylode.html: $(build_dir)/%.owl.ttl | $(build_dir)/
	$(venv_run) pylode -i $(<) -p ontdoc --css true --getcuriesonline true --outputfile $(@) -f html
	@echo google-chrome $(@)

########################################################################
# toolchain
########################################################################

$(owl2vowl_jar): | $(tools_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	curl http://downloads.visualdataweb.de/owl2vowl_0.3.7.zip | bsdtar -xvf - -C '$(tools_dir)' owl2vowl.jar

$(build_dir)/webvowl/index.html: | $(build_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	curl http://downloads.visualdataweb.de/webvowl_1.1.7.zip | bsdtar -xvf - -C "$$(dirname -- "$(@)")"

$(rdflint_jar): | $(tools_dir)/
	mkdir -vp $$(dirname -- "$(@)")
	curl -L -o '$(@)' https://jitpack.io/com/github/imas/rdflint/0.1.4/rdflint-0.1.4.jar

$(robot_jar): | $(tools_dir)/
	curl -L -o '$(@)' https://github.com/ontodev/robot/releases/latest/download/robot.jar

rdflint: $(rdflint_jar)
	$(rdflint)

jena_url=https://archive.apache.org/dist/jena/binaries/apache-jena-4.0.0.tar.gz
$(jena_riot): | $(tools_dir)/ $(cache_dir)/ $(jena_dir)/
	cd $(cache_dir) && wget -c "$(jena_url)"
	tar --strip-components=1 -zxvf "$(cache_dir)/$(notdir $(jena_url))" -C '$(jena_dir)'

$(widoco_jar): | $(tools_dir)/
	curl -L -o $(@) https://github.com/dgarijo/Widoco/releases/download/v1.4.14/widoco-1.4.14-jar-with-dependencies.jar

vocbench_dir=$(tools_dir)/vocbench3
vocbench_start=$(vocbench_dir)/bin/start
vocbench_url=https://bitbucket.org/art-uniroma2/vocbench3/downloads/vocbench3-8.0.0-full.zip
$(vocbench_start): | $(tools_dir)/ $(cache_dir)/ $(vocbench_dir)/
	cd $(cache_dir) && wget -c $(vocbench_url)
	bsdtar --strip-components 1 -xvf $(tools_dir)/$(notdir $(vocbench_url)) -C $(vocbench_dir)
	chmod -R +x $(vocbench_dir)/bin

$(ontoviz): | $(tools_dir)/
	mkdir -vp $(dir $(@))
	curl -L https://github.com/usc-isi-i2/ontology-visualization/archive/master.tar.gz | bsdtar -xvf - -C $(dir $(@)) --strip-components 1

.PHONY: vocbench
vocbench: $(vocbench_start)

.PHONY: distclean
distclean: clean-$(tools_dir)/

toolchain: $(jean_dir) $(rdflint_jar) $(robot_jar) $(jena_riot) $(tools_dir)/owl2vowl.jar $(build_dir)/webvowl/index.html $(ontoviz) $(widoco_jar)
	#pip3 install --user --upgrade pipenv poetry
	# dnf install
	python -m pip install --user --upgrade poetry
	poetry install
	#pipenv install
	npm install -g jsonld-cli live-server
	python -m pip install --user --upgrade xdot
