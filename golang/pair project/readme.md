# 1. cara membuat branch dan langsung ke branchnya

    git switch -c feature-sql 

# 2. cek branch

    git branch 

    ** jika sudah sesuai masuk ke feature-sql maka push

# 3. git push branch

    git push -u origin feature-sql

    git push


# Ambil perubahan terbaru dari main
git switch main
git pull origin main

# Buat branch baru dan pindah ke branch tersebut
git switch -c feature/student

# Setelah selesai coding
git add .
git commit -m "Menambahkan fitur student"

# Push branch ke GitHub
git push -u origin feature/student