# Extract-Bundle

## Overview

The extract-bundle tool is an **experimental** analytic tool which uses the Operator Framework solutions. Its purpose is to obtain and report all operator bundles from an index catalog image.

## Pre-requirements

- go 1.16 (only if you would like to install from the source)
- docker or podman

## Generating the reports

Now, you can audit all operator bundles of an image catalog with: 

```sh 
extract-bundle extractBundle --index-image=registry.redhat.io/redhat/redhat-operator-index:v4.9
```

## How extract-bundle tool works? 

- Extract the database from the index-image provided
- Perform SQL queries to obtain the data from the index db
- Download and extract all bundles name by using the operator bundle path which is stored in the index db  
- Get the required data(bundle paths and bundle names) for the report from the operator bundle manifest files 
