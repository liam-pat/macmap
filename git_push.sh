#!/bin/bash

check_to_commit(){
   if ! sha256sum -c sha256sum.txt; then
          str1="v1.0."
          current_v=`git ls-remote --tags origin | grep -o '[0-9]$' | tail -1`
          git_tag=`expr $current_v + 1`
          tag_str=$str1$git_tag

          echo "current tag:" $current_v
          echo "new tag:" $tag_str

          sha256sum MAM.csv MAS.csv MAL.csv > sha256sum.txt
          git status

          git config user.name 'github CI/CD'
          git config user.email 'github-actions@github.com'
          git add sha256sum.txt MAM.csv MAS.csv MAL.csv
          git commit -m "DB-update: Auto update the mac db"
          git push

          git tag -a $tag_str -m "New tag : mac db updating"
          git push origin $tag_str
      fi
}

main() {
    check_to_commit
}
main