cd src
src.exe --files=data/csp*.cnv
ncdump output/OS_CASSIOPEE_CTD.nc > output/FichierResult.txt
notepad++.exe output/FichierResult.txt