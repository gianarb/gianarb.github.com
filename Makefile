.PHONY: init tail
init:
	@docker run --rm -it --name gianarb_blog -p 80:4000 -v ${PWD}:/srv/jekyll jekyll/jekyll jekyll serve -w --incremental --drafts --unpublished
tail:
	@docker logs -f gianarb_blog
