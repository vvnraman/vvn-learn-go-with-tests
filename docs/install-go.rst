**********
Install go
**********

Official install binaries
=========================

Official binaries available at https://golang.org/doc/install

- The download URL for linux is https://golang.org/dl/go1.16.4.linux-amd64.tar.gz

- Given the version, here ``1.16.4``, my custom install steps which can ideally
  be put into a script to be executed

  .. code-block:: sh

     GO="go"
     GO_VERSION="1.16.4"
     GO_TAR_GZ="${GO}${GO_VERSION}.linux-amd64.tar.gz"
     GO_TAR_GZ_URL="https://golang.org/dl/${GO_TAR_GZ}"
     mkdir -p "${HOME}/code/tools/golang"
     GO_TAR_GZ_PATH="${HOME}/code/tools/golang/${GO_TAR_GZ}"
     curl --create-dirs --fail --location --output "${GO_TAR_GZ_PATH}" "${GO_TAR_GZ_URL}"
     tar --directory "${HOME}/code/tools/golang" --extract --gzip --file "${GO_TAR_GZ_PATH}"
     sudo mv "${HOME}/code/tools/golang/go" "/usr/local/go-${GO_VERSION}"
     sudo rm /usr/local/go
     sudo ln -s /usr/local/go-${GO_VERSION} /usr/local/go

     # one time only
     export PATH=$PATH:/usr/local/go/bin

  What it does

  - Forms the filename from the version string ``go1.16.4.linux-amd64.tar.gz``

  - Forms the URL by prefixing ``https://golang.org/dl/`` to  the filename

  - Sets up a download directory for the install package at
    ``"${HOME}/code/tools/golang"``

  - Uses :program:`curl` to download the install package into the download
    directory, and exacts it there.

  - Moves the extracted archive to ``/usr/local/go-${GO_VERSION}`` and symlinks 
    ``/usr/local/go`` to the the versioned directory.

  - One time only - Add ``/usr/local/go/bin`` to path. Subsequent install of
    newer golang releases don't need this.

- A more recent version of the install routing might be present in my Notebook.

Editor setup - vim
==================

Install https://github.com/fatih/vim-go plugin

.. code-block:: vim

   Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }

``GoUpdateBinaries`` takes a while.

Project setup
=============

.. code-block:: sh

   mkdir -p ~/code/langs/go/learn-go-with-tests/learn-go-with-tests
   cd ~/code/langs/go/learn-go-with-tests/learn-go-with-tests
   go mod init learn-go-with-tests

This gave us the following ``go.mod`` file

.. code-block:: go

   module learn-go-with-tests

   go 1.16

