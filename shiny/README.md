# Getting started with Shiny (using R)

This sample demonstrates using the Google Genomics API from R along with the
[Shiny web application framework](http://shiny.rstudio.com/).

For more basic getting started with R examples, see the
[Bioconductor GoogleGenomics repository](https://github.com/Bioconductor/GoogleGenomics).

You can view a hosted version of this Shiny app here:

  http://googlegenomics.shinyapps.io/getting-started

To run the Shiny app locally, follow the instructions below:

1. If you have not already done so, follow the Google Genomics
   [sign up instructions](https://cloud.google.com/genomics/install-genomics-tools#authenticate)
   to generate and download a valid ``client_secrets.json`` file.

   Save the "client ID" and "client secret" values from the credential.

2. You can run the code in this sample directly from github or from your
   local file system.

   To run this code directly from github. issue the following in R:

  ```
  devtools::install_github("shiny", "rstudio")
  library(shiny)
  runGitHub("getting-started", "googlegenomics", subdir = "R")
  ```

   To run this code from your local copy, issue the following in R:

  ```
  devtools::install_github("shiny", "rstudio")
  library(shiny)
  runApp('path_to_local_directory')
  ```

  where ``path_to_local_directory`` points to the location with the server.R
  and ui.R files. If this is already the working directory for your R
  instance, then the path parameter can be omitted.

3. The ``runGitHub`` or ``runApp`` command will start a Shiny server
   on your machine and open a local web port. You should see a message
   such as:

  ```
  Listening on http://127.0.0.1:6652
  ```

   The command should also open your web browser to the listed local URL.

4. On the displayed web page, fill in the client ID and secret values.

5. If you have not already authenticated and cached your credentials, your
   browser should now prompt you to authorize access to the Genomics API.
   On successful authorization, the requested genomic data should be
   fetched and displayed in the Shiny app.
   in the 
