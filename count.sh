so1=$(find . | wc -l)
ket_qua=$(((so1) * 5))
# printf " \t\vTotal files * 5: $ket_qua\v\n"
printf "\t\vTotal files * 5: %d\v\n" "$ket_qua"
