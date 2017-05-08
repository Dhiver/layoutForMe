# Layout For Me

----

LayoutForMe is an open source document generator that helps you to focus more on the content.

You write [Markdown](https://daringfireball.net/projects/markdown/), it gives you a PDF document (for now...).

----

## Features

* Retrieve history of your document throught git tags
* Templating templates !
* Template translation
* Template inclusion

## Dependencies

* Go
* A LaTeX distribution (in order to have the `lualatex` command)
* libgit2
* pandoc

## Install

`go get github.com/Dhiver/layoutForMe`

`go install github.com/Dhiver/layoutForMe`

## Quick Start

You can see an example workdir [HERE](https://github.com/Dhiver/layoutForMe_example).

## Quick Start Explained

### Create a workdir

In a *new* git repository, put your markdown files inside a folder named `content`. 

Ex.
```
.
|- content/
|  |- intro.md
|  |- conclusion.md
```

Create an empty `build` folder and an empty file named `configuration.yaml`.

### Give metadata about your document

In a file called `metadata.yaml` set metadata about your document:

Here is an example :

```yaml
lang: fr-FR
dateLayout: "02/01/2006"

title: Bonjour à tous !
author:
    - John Smith
    - Foo Bar

abstract:
date:
version:

includeOrder:
    - intro.md
    - conclusion.md

output:
    - name: mon_document
      extention: pdf
      template: article.tex.tmpl
```

If you leave the `abstract` field blank, the content of a file nammed `README.md` will be load into.

If the `date` field is leave empty, the current date will be automatically inserted.

If the `version` field is empty, layoutForMe will look for tags in the current git repository and insert then as entries.

### Create a template

A template will be a LaTeX document (.tex) with [Go templating language](https://golang.org/pkg/text/template/).

[Here is an example](https://github.com/Dhiver/layoutForMe_example/blob/master/templates/article.tex.tmpl)

In the template, you can use variables defined in [`metadata.go`](metadata.go) as well as functions defined in [`template.go`](template.go#L30) (near funcMap).

### Layout !

`layoutForMe -path /your/document/path`

The generated document will be in the `build` folder.

#### Configuration

In the configuration file : `configuration.yaml` you can change some values :

* `metaFile` (Default: metadata.yaml)
* `buildFolder` (Default: build)
* `textFolder` (Default: content)
* `pictureFolder` (Default: img)
* `templateFolder` (Default: templates)
* `translateFolder` (Default: translations)
* `latexEngine` (Default: lualatex)
