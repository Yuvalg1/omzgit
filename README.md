## omzgit

I had a problem I was using almost the exact same git commands every single time. I used the ```cli```, though excellent I was only using a handful of commands. Sometimes, I even used the ```VSCode``` UI where I was only opening it to commit, stage and look at the git graph nicely.

I was looking for other options, and a friend of mine told me about [ohmyzsh](https://github.com/ohmyzsh/ohmyzsh). He told me it has a git cheat sheet exactly for my requirements. Around the same time I also found out about [lazygit](https://github.com/jesseduffield/lazygit). I started diving into list of available options at my disposal and it was great, although a bit daunting. At the time I was only using a handful of commands, and I learned a thing or two about designing a tiny bit of UI, and I thought to myself why not create a project much smaller than either of them, fit exactly to my needs.

I scanned the cheat sheet trying to find patterns and I wanted to create the program to fit it as best as I could. Then I had the realization that I can use pages. Then, it fits great and I can reuse the same letters. Separate it into multiple pages for better readability, and give a visual feedback to each letter I type.

And that's how the idea was born.

## Installation

### Releases

Simply download the latest available binaries and run them. In [Releases](https://github.com/Yuvalg1/omzgit/releases).

### Manual

Please [install go](https://go.dev/doc/install) before continuing.
```
git clone https://github.com/Yuvalg1/omzgit.git
cd omzgit
go install
```

> [!NOTE]
> I'm working on more installation options.

## Usage

To enter ```omzgit```, call it in your terminal inside a git repository:

```go
$ omzgit
```

In any page, the following behavior is consistent:

```f``` - fetch changes

```l``` - pull changes from origin

```p``` - push changes to origin

```q``` or ```ctrl+c``` - quit

```/``` - open search filter

```y``` - yanks relevant information from current row

### Files

Here you can stage, unstage, discard, and commit files. On the left are the files changed, and on the right is the diff.

https://github.com/user-attachments/assets/7cfd2b71-78af-4454-bb46-2233ae79563c

### Branches

Here you can checkout, delete, merge, rebase, and switch branches. On the left are the branch names, and on the right of the current row is the behind|ahead the the last changed time difference.

https://github.com/user-attachments/assets/19af9335-b887-4494-a267-f4f00da9fa5c

### Commits

Here you can checkout, reset and look through all of the commits.

https://github.com/user-attachments/assets/3e31cedf-e81b-46e4-9ceb-bd5740980c19
