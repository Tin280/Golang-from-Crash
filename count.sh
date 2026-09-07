so1=$(ls -lR| wc -l)
so2=$(find . -type f -o -type d |wc -l)
ket_qua=$(((so1) * 5))
printf "&#11;Total files * 5: $ket_qua$&#11"
printf  "$(ls -lR)"