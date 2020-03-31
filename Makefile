######################################################
ORG=ucsd-cse223b-sp20
ASSN=lab1
STARTER=lab1-starter
######################################################

check: 
	cd ${CURDIR}/src/triblab/ && $(MAKE) all
	cd ${CURDIR}/src/triblab/ && $(MAKE) testv-lab1

turnin: 
	git commit -a -m "turnin"
	git push origin master

upstream:
	git remote add upstream git@github.com:$(ORG)/$(STARTER).git

update:
	git pull upstream master