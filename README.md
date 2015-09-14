# Golang: vendor's using example

#### Use vendor packages
[reference](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)   
In golang version 1.5, you have to set vendoring by manual.   
* Do   
`$ export GO15VENDOREXPERIMENT=1`
* Put your package under `vendor` folder, path struct must be placed the same as `$GOPATH/src`.
* Import packages as usual, golang will find out packages by itself.
