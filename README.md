[kube-tmuxp](https://github.com/thecasualcoder/kube-tmuxp/) has been migrated to Golang. I would recommend using that, as it's well 
maintained and is robust with tests!

This project is currently not maintained.

# kube-tmuxp-config-gen

This project has been inspired from [kube-tmuxp](https://github.com/arunvelsriram/kube-tmuxp/) and is a Golang port of it

## Why develop a Golang port?

When I tried [kube-tmuxp](https://github.com/arunvelsriram/kube-tmuxp/), it was quite tough to install it itself. I had Python 3.7 version and I wasn't able to make it work. Then I had to install Python 3.6 as shown in the repository README to make it work. It just wouldn't work with even a minor version change in Python. This and the conversation with my colleagues lead me to writing this in Golang. By writing it in Golang, I just have to compile it once for a given architecture and then give the single static binary and the binary would just work. That's what I am trying to achieve here 😄

## Learnings

I have shared my [learnings here](LEARNINGS.md)
