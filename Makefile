.PHONY: init tail
init:
	@docker run --rm --name gianarb_blog -p 80:4000 -v ${PWD}:/opt/site gianarb/jekyll jekyll serve -w --incremental
tail:
	@docker logs -f gianarb_blog
