# custom_firefox_search
run local webserver to host OpenSearch XML files to add to Firefox as custom search engines.

looks in specified directory (default is current directory) for XML files containing OpenSearch definitions. 
will then create small webpage with <link> in <head> to allow Firefox to add these search engines.
see [wolfram.xml](engines/wolfram.xml) for example. for more details on how to create the XML files, see https://developer.mozilla.org/en-US/docs/Web/OpenSearch

main use for this is as an alternative to firefox addon https://addons.mozilla.org/en-US/firefox/addon/add-custom-search-engine/ to prevent uploading anything to file.io.
