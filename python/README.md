# Getting started in Python

1. If you have not already done so, follow the Google Genomics [sign up instructions](https://cloud.google.com/genomics/install-genomics-tools#authenticate) to generate and download a valid ``client_secrets.json`` file.  

2. Copy the client_secrets.json file into this directory.

3. [Install pip](http://www.pip-installer.org/en/latest/installing.html)

4. (Optional) Create and use a [Python virtualenv](http://docs.python-guide.org/en/latest/dev/virtualenvs/)

    ```
    pip install virtualenv
    virtualenv python_sample
   
    source python_sample/bin/activate
    ```

5. Install the python client library and run the code:

    ```
    pip install --upgrade google-api-python-client
    python main.py
    ```

6. To exit the virtualenv (if used):

    ```
    deactivate
    ```

# More information

* [Google Genomics client library](https://cloud.google.com/genomics/v1/libraries)
* [Pydoc reference for the Genomics API](https://developers.google.com/resources/api-libraries/documentation/genomics/v1/python/latest/)

# Troubleshooting

[File an issue](https://github.com/googlegenomics/getting-started-with-the-api/issues/new)
if you run into problems and we'll do our best to help!
