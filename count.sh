# so1=$(ls -R | wc -l)
so1 =$(find . -mindepth -type d | wc -l)
ket_qua=$((so1 * 5))
echo "&#11;Total files * 5: $ket_qua$&#11"
