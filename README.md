# kf

A tool to decode [baldr](https://github.com/uswitch/baldr) files.

## Usage

    $ bdcat --help
    usage: bdcat [<file>]

    Flags:
      --help  Show help.

    Args:
      [<file>]  baldr file to decode

## Example

Let's imagine we have a compressed baldr file on S3 which contain JSON records, each with a timestamp.
This pipeline can be used to get all the timestamps:

    $ s3cmd get s3://bucket/file.baldr.gz | zcat | bdcat | jq .timestamp

## Prerequisites

* Go

## Installation

    $ go get -u github.com/uswitch/bdcat
