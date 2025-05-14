#!/bin/bash

check_to_commit(){
   if ! sha256sum -c sha256sum.txt; then
          echo "Detected change in CSV checksum. Proceeding to commit."

          latest_tag=$(git tag | grep '^v1\.0\.' | sort -V | tail -1 || echo "v1.0.0")
          current_v=${latest_tag##v1.0.}
          next_v=$((current_v + 1))
          tag_str="v1.0.$next_v"

          echo "Current tag: $latest_tag"
          echo "New tag: $tag_str"

          sha256sum MAM.csv MAS.csv MAL.csv > sha256sum.txt

          git config user.name 'github CI/CD'
          git config user.email 'github-actions@github.com'
          git add sha256sum.txt MAM.csv MAS.csv MAL.csv
          git commit -m "DB-update: Auto update the mac db (v1.0.$next_v)"
          git push

          git tag -a "$tag_str" -m "New tag: mac db updating"
          git push origin "$tag_str"
   else
        echo "No changes detected in CSV files."
   fi
}

main() {
    check_to_commit
}
main