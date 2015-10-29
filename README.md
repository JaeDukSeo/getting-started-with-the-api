getting-started
===============

The repository contains examples of how to integrate with Google Genomics API.
The code within each language-specific folder demonstrates the same things:

* Getting the read bases for NA12872 at a specific position
* Getting the variant overlapping that same position, and outputting the genotype

Prerequisites
-------------

1. Clone or fork this repository.

1. All languages will require a valid Client ID and Client secret.
   If you have not already done so, follow the Google Genomics
   [sign up instructions](https://cloud.google.com/genomics/install-genomics-tools#authenticate)
   to generate and download a valid ``client_secrets.json`` file.
