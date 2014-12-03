

main = do print $ length [a| a<-[1..1000], n<-[1..100], (length (show (a^n))) == n]
