package templates

// README is the contents of the README.md file.
const README = `# dslink-%s
## %s DSLink

A DSLink for <complete this>

`

// DSLinkJSON is the contents of dslink.json file.
const DSLinkJSON = `{
  "name" : "dslink-dart-%s",
  "version" : "1.0.0",
  "description" : "%s DSLink",
  "main" : "bin/run.dart",
  "engines" : {
    "dart" : ">1.13.0"
  },
  "getDependencies" : [
    "pub get"
  ],
  "configs" : {
    "broker" : {
      "type" : "url"
    },
    "name" : {
      "type" : "string",
      "default" : "%[2]s"
    },
    "token" : {
      "type" : "string"
    }
  }
}

`

// PubSpec is the contents of the pubspec.yam file.
const PubSpec = `name: dslink_%s
description: %s DSLink
environment:
    sdk: ">=1.13.0 <2.0.0"
dependencies:
    dslink:
        git: https://github.com/IOT-DSA/sdk-dslink-dart.git

`

// DartRun is the template for the bin/run.dart file.
const DartRun = `import "package:dslink/dslink.dart";

import "package:dslink_%s/%[1]s.dart";

main(List<String> args) async {
    LinkProvider link;

    link = new LinkProvider(args, "%s", autoInitialize: false, profiles: {

    }, defaultNodes: {

    });

    link.init();
    await link.connect();
}

`
