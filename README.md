# gocr
OCR pipeline in golang


## Configuring Tesseract

This script uses the Tesseract module. In order to configure the right language follow the steps below.

### installed languages
To get an overview of installed languages

```
tesseract --list-langs
```

Missing languages need to be installed. Using the example commandline below, the dutch language is installed (`nld`)

```
sudo apt-get install tesseract-ocr-nld
```
