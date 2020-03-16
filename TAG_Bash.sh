#!/bin/bash


:'
cria um arquivo de backup com o endereco digitado pelo usuario
no segundo argumento e o diretorio a ser copiado eh fornecido
no primeiro argumento, que tambem da nome ao arquivo de backup
'
bk () {
	tar -cpzf $2/`date +%Y-%m-%d`_$1.tar $1
}

bk $1 $2

:'
nesta funcao, eh criada uma variavel que contem o dia da execucao
do script e de tres anteriores, deletando todos os arquivos com
a extensao .tar e que passaram de tres dias.
obs: para que todos os arquivos de backup "fora de validade"
sejam apagados de uma vez, eh necessario que estejam na mesma
pasta da execucao do script
'
deletar () {
	local s
	for dia in {0..3}; do
		s=$s$(date -I -d "today -$dia days")'|'	
		#ex de output: 2019-03-07|2019-03-06|2019-03-05|2019-03-04|
	done 

	ls | grep .tar | grep -v -E ${s%?} | xargs rm
	#${s%?} para retirar o ultimo '|', senao tal grep nao funcionaria corretamente
}

deletar
