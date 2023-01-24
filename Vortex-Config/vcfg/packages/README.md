# Vortex-Config

**How to add a command to ```vcfg```?**

Pretend we are adding a command called ```examplecommand``` such that we can call ```vcfg examplecommand```


**Create the command itself  ```cmds/examplecommand.go```**

This is the command that calls its implementation from ```vortexcfg/Examplecommand.go``` and is called by ```main.go```


**Create the implementation of the command ```vortexcfg/Examplecommand.go```**

This is the function that is called by the implementation inside of ```cmds/```
