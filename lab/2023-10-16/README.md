# 1 Ordinamento per inserimento
Scrivete una funzione che legga da standard input una sequenza di interi distinti terminati da
0, memorizzandoli in un vettore ordinato (valutate se è più opportuno usare un array o una
slice): ogni volta che viene letto un nuovo intero, il vettore viene scorso fino a trovare l’esatta
collocazione del numero, quindi si crea lo spazio per il nuovo numero spostando in avanti i
numeri successivi già memorizzati.
Questo algoritmo è utile per riempire un vettore mantenendolo ordinato ad ogni passo.
# 2 Ordinament per selezione
Scrivete una funzione che riceva una slice di interi e la ordini usando l’algoritmo SelectionSort:
alla fine dell’esecuzione, la slice originaria passata come argomento dovrà risultare ordinata. Può
essere utile restituire la stessa slice (ordinata), ad esempio per poterla passare come argomento
ad altre funzioni, come in fmt.Println(selectionSort(v)).


Versione iterativa La funzione selectionSortIter( int a[]) ripete la seguente operazioni
tante volte quanto è lunga la slice da ordinare: per ogni prefisso di lunghezza n (con n inizialmente pari alla lunghezza della slice) cerca nel prefisso l’elemento massimo e lo scambia con
quello nell’ultima posizione del prefisso

Versione ricorsiva Scrivete una versione ricorsiva dello stesso algoritmo: la funzione selectionSortRec
deve funzionare come segue:
- innanzitutto cerca nel vettore l’elemento massimo e lo sposta nell’ultima posizione;
- poi richiama se stessa ricorsivamente per ordinare i primi n−1 elementi del vettore.
La base della ricorsione è data dai vettori di lunghezza 0 o 1, che sono sempre ordinati.
