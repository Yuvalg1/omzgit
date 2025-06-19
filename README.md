## omzgit

A __GUI__ app is like cruising with a fancy dashboard and auto-pilot. It looks great and gets you there with minimal effort.

The __CLI__? That’s like riding a manual motorcycle: raw, fast, and thrilling, but only if you know what you’re doing.

what about __omzgit?__

Not too long ago, I found out about [lazygit](https://github.com/jesseduffield/lazygit) and [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh). Both projects are amazing and huge inspirations for this project. 

lazygit is great, but I wanted it to look more like what I am used to seeing in VSCode git editor. VSCode is good, but when I started to learn Vim I wanted the functionality of shortucts and the feeling of not using a mouse. So when I looked at ohmyzsh, specifically about the git part documented in the git [cheat sheet](https://kapeli.com/cheat_sheets/Oh-My-Zsh_Git.docset/Contents/Resources/Documents/index), I had a vision.

What if I make a terminal UI project which looks simplistic for a newbie like me, but holds the power of commands inspired by the Oh-My-Zsh Git Cheat Sheet? And that's when omzgit was born.

## Installation

### Manual

Please [install go](https://go.dev/doc/install) before continuing.
```
git clone https://github.com/Yuvalg1/omzgit.git
cd omzgit
go install
```

> [!NOTE]
> I'm working on more installation options.

## Run omzgit

```
omzgit
```

## Usage

```f``` - fetch changes

```l``` - pull changes from origin

```p``` - push changes to origin

```q``` or ```ctrl+c``` - quit

### Pages

There are currently two branches, ```Files``` and ```Branches```. (Stay tuned for Commits)

To switch between pages, you use the prefix ```g```, meaning ```git``` or ```goto```.

```gb``` - for ```Branches``` page.

```gf``` - for ```Files``` page.

This also refreshes the data in the page. If you want an easier way to refresh page data, simply press ```esc```.

### Branches

To move between branches, use the arrow keys or ```j/k``` (more on that later)

```b``` - create a new branch

```c``` - checkout to current selected branch

```d``` - delete a branch 

```D``` - force delete a branch

### Files

```a``` - stage selected file

```A``` - stage all files

```c``` - open commit popup

```d``` - discard selected file

```D``` - discard all files

### Commit Popup

```F``` - take a commit message from the file name

```m``` - enter a commit message

```o``` - toggle options menu

#### Commit Options

```a``` - --amend

```e``` - --edit

```E``` - --no-edit

```n``` - --no-verify

```y``` - --allow-empty
