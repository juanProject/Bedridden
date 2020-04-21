# Projet Go

Programme permettant de gérer de manière concurrente les lits d'un hôpital.
 
Un médecin, lorsqu'il hospitalise un malade, décrémente de 1 le nombre de lits dans cet hôpital.
L'idée derrière ce programme est d'avoir un fichier d'entrée contenant le nombre de lits pris et celui des lits libres, et,
lancer deux (ou plus) programmes Go concurrents, qui auront cet input et qui communiqueront entre eux.

Par exemple, le médecin du programme A va pouvoir prendre un lit et son programme communiquera au programme B le fait 
que l'on a décrémenté de 1 le nombre de lits, et pareil lorsqu'un lit se libère.

# Getting Started

build project :

```shell script
go build
```

To change default params go change data.json file 

Example : 
```json
{
  "TotalBeds": 1000, 
  "NaBeds": 447
}
```

# Ressources

[code-lab](https://blog.golang.org/codelab-share)

[processes communicate with shared memory in golang](https://gist.github.com/jedy/3282764)

[shm](https://godoc.org/github.com/ghetzel/shmtool/shm)