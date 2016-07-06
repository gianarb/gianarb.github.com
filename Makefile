.PHONY: init tail
init:
	@docker run -d --name gianarb_blog -p 4000:4000 -v ${PWD}:/opt/site gianarb/jekyll
tail:
	@docker logs -f gianarb_blog
