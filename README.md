# Creates a table of contents

<img src=http://s.natalian.org/2016-06-05/toc.gif alt="Demo">

An element with `data-fill-with="table-of-contents"` is filled with an ordered list of the headers.

## Install

	go get github.com/kaihendry/toc

## Example

`foo.html` contains:

	<h1>FAQ</h1>
	<nav data-fill-with="table-of-contents" id="toc"></nav>
	<h3 id="how-do-i-create-a-faq">How do I create a FAQ?</h3><p>Like this!</p>

Run the tool over the HTML

	$ toc foo.html

And the following with print to stdout

	<html><head></head><body><h1>FAQ</h1>

	<nav data-fill-with="table-of-contents" id="toc"><ol>
	<li><a href="#how-do-i-create-a-faq">How do I create a FAQ?</a></li>
	</ol></nav>

	<h3 id="how-do-i-create-a-faq">How do I create a FAQ?</h3><p>Like this!</p>
	</body></html>
