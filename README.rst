Practice - Learn Go with Tests
##############################

My **Learn go with Tests** exercises.

- https://quii.gitbook.io/learn-go-with-tests/

****
Docs
****

Published at https://vvnraman.github.io/vvn-learn-go-with-tests/index.html

- Generated using sphinx
- Managed via poetry

Generate docs locally
=====================

.. code-block:: sh

   make docs

- Runs sphinx ``Makefile``

  - via ``poetry run``

    - via top level Makefile

Publish docs
============

.. code-block:: sh

   make publish

----

One-time only
*************

Docs setup
==========

- Create a git repo with a ``.gitignore``

  .. code-block:: sh

     echo "docs/_build" >> .gitignore
     git init .
     git add .gitignore
     git commit -m "Initial commit"

- Generate poetry boilerplate

  .. code-block:: sh

     poetry init
     poetry add sphinx
     poetry add sphinx-rtd-theme
     poetry add sphinxcontrib-seqdiag
     poetry add sphinxcontrib-nwdiag
     poetry add sphinxcontrib-blockdiag
     poetry add sphinxemoji
     poetry add sphinx-tabs

- Setup virtualenv

  .. code-block:: sh

     poetry install

- Create sphinx documentation boilerplate

  .. code-block:: sh

     poetry run sphinx-quickstart

- Update ``docs/conf.py`` ``extensions``

  .. code-block:: python

     extensions = [
             "sphinxcontrib.seqdiag",
             "sphinxcontrib.nwdiag",
             "sphinxcontrib.blockdiag",
             "sphinxemoji.sphinxemoji",
             "sphinx_tabs.tabs"
     ]

- Update ``docs/conf.py`` ``html_theme``

  .. code-block:: python

     html_theme = 'sphinx_rtd_theme'
