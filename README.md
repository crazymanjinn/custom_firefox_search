# custom_firefox_search
run local webserver to host OpenSearch XML files to add to Firefox as custom search engines.

looks in specified directory (default is current directory) for XML files containing OpenSearch definitions. 
Will then create small webpage with <link> in <head> to allow Firefox to add these search engines.
See [wolfram.xml](engines/wolfram.xml) for example.

for more details, see https://developer.mozilla.org/en-US/docs/Web/OpenSearch
