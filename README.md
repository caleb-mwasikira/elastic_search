Hi @all I need some help here


I need to perform a search on a .txt file with over 1M rows.
The file gets updates very frequently, updates may be a post, delete, update etc..

The crazy part is that due to this frequent changes I have to reread the file on every query.

The whole operation should take a maximum of 40ms


I have tried MMAP and Grep strategy in Linux but none can do.

I  will appreciate any kind of help. The file is not sorted. I'm looking for a solution with minimum overhead time for file processing