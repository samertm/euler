
import qualified Data.Set as Set

main :: IO ()
main = do print $ Set.size $ Set.fromList([a^b | a <-[2..100], b<-[2..100]])
