# DataStager
A simple application to stage data from and to Kubernetes pods
Requires S3 config to be present in the pod, these can be mounted via environment variables. The application uses the AWS S3 SDK for Golang to perform file up and downloads. The details for the configuration can be found here: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html. A custom endpoint can be set with a CLI argument.

## Download example
Two objects should be downloaded, both are in the bucket stagingexample with the keys examples/examplefile1.json and example/examplesfile2.json.
The data should be uploaded under the key examples.
The data is stored in a self deployed s3 object storage with the endpoint s3.example.com.

```bash
datastager download -b stagingexample, stagingexample -k example/examplesfile1.json, example/examplesfile2.json -d /home/example/data -e s3.example.com
```

## Upload example
Two objects should be staged, both are in the bucket stagingexample with the keys examples/examplefile1.json and example/examplesfile2.json.
The data should be stored in the directory /home/example/data.
The data is stored in a self deployed s3 object storage with the endpoint s3.example.com.

```bash
datastager upload -e s3.computational.bio.uni-giessen.de -k test/data/1 -b BaktaTest -f /home/marius/Data/SmallTestData/test1.txt,/home/marius/Data/SmallTestData/test2.txt
```