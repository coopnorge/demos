all: default
default:

# .ONESHELL:
SHELL=bash
# .SHELLFLAGS=-ec -o pipefail

datadir=share
localstatedir=var
statesdir=$(datadir)/states
demodir=$(localstatedir)/demo


clean: clean-$(localstatedir)/

# setup: | $(demodir)/
# 	cp -rv $(statesdir)/* $(demodir)

$(demodir)/%: $(statesdir)/%
	mkdir -vp $(demodir)
	rsync -av --delete $(<)/ $(@)/


# demo_000_dir:
# 	rsync

pause=read -s -n 1

demo_000_dir=000-nodigest
demo_010_dir=010-olddigest
demo_020_dir=020-locked

setup: \
	$(demodir)/$(demo_000_dir) \
	$(demodir)/$(demo_010_dir) \
	$(demodir)/$(demo_020_dir) \

.PHONY: clean-%/
clean-%/:
	@{ test -d $(*) && { set -x; rm -vr $(*); set +x; } } || echo "directory $(*) does not exist ... nothing to clean"


.PRECIOUS: %/
%/:
	mkdir -vp $(@)
