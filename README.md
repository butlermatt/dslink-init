# DSlink-Init

DSLink-Init is a program to generate scaffolding for DSA DSLink projects. This
tool is still very much in early development. At the moment it only supports
generating the scaffolding for Dart based projects.

## Usage

```
dslink-init <projectName> <nodeName>
```

`projectName` is the name for the project sources. This will automatically
create a project with the name *dslink_projectname* and generate several files
by the name of *projectname.dart*.

`nodeName` is the default name for the nodes which will appear in DSA/DGLux for
this project.

This tool will generate the following files:

```
README.md
dslink.json
pubspec.yaml
bin\
bin\run.dart
lib\
lib\models.dart
lib\<projectName>.dart
lib\src\
lib\src\nodes\
lib\src\models\
```
