# Clone of a single folder from a Git Repo

Checkout jednoho adresáře v repozitáři (https://github.com/JakubMatejka/kac-test/tree/main/template1) se v gitu udělá takto:

```
git clone --depth 1 --filter=blob:none --no-checkout https://github.com/JakubMatejka/kac-test.git
git checkout main -- template1
```

Implementace checkoutu v Go-Git (https://github.com/go-git/go-git/blob/e4fcd078d42e945581616855ab78d8b7ed12df6c/worktree.go#L153) ale neumí filtrovat paths (chybí to v `CheckoutOptions`: https://github.com/go-git/go-git/blob/e4fcd078d42e945581616855ab78d8b7ed12df6c/options.go#L277)


Vycházel jsem z in-memory příkladu: https://github.com/go-git/go-git#in-memory-example. 

- v [clone-all.go](./clone-all.go) je obyčejný vyklonování repa
- v [clone-override.go](./clone-override.go) je totéž ale s lokálně přetíženým `Storer` (https://github.com/go-git/go-git/blob/master/storage/memory/storage.go) a `Filesystem` (https://github.com/go-git/go-billy/blob/master/memfs/storage.go). Obě jsem se pokoušel přetížit a do různých míst strčit filtrování na konkrétní adresář. (Naposledy sem do `Write` metody: https://github.com/JakubMatejka/kac-test/tree/main/memory.go#L283 - resp. https://github.com/go-git/go-billy/blob/7ab80d7c013de28ffbb1ca64b9bbf8dd1cbd81c5/memfs/memory.go#L279).
- v [clone-all-commit.go](./clone-all-commit.go) pokus s vylistováním souborů z commitu (asi k ničemu :-)
- v [clone-git2go.go](./clone-git2go.go) je funkční implementace s git2go a přímým checkoutem vybranýho adresáře na disk
