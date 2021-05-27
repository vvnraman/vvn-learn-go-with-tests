Practice - Learn Go with Tests
##############################

My **Learn go with Tests** exercises.

- https://quii.gitbook.io/learn-go-with-tests/

****
Docs
****

- Generated using sphinx
- Managed via poetry

Generate docs
=============

- Run sphinx via poetry

  .. code-block:: sh

     poetry run make -C docs html

One-time only
=============

- Generate initial boilerplate

  .. code-block:: sh

     poetry init
     poetry add sphinx
     poetry add sphinx-rtd-theme
     poetry add sphinxcontrib-seqdiag
     poetry add sphinxcontrib-nwdiag
     poetry add sphinxcontrib-blockdiag
     poetry add sphinxemoji
     poetry add sphinx-tabs
     poetry install
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
