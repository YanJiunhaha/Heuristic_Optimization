set term png enhanced font 'Verdana,10'
set output 'graph.png'

set isosamples 30,30
set hidden3d
#set contour
#set cntrparam levels 20

set title "hw01"
set xlabel "X"
set ylabel "Y"
set xrange [-5:5]
set yrange [-5:5]

splot 4/((x-2)**2+(y-2)**2+1)+3/((x-2)**2+(y+2)**2+1)+2/((x+2)**2+(y-2)**2+1)
