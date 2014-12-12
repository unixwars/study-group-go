import Data.List

main = interact (show.(map count).group.sort.words.(filter (not.punctuation)))

punctuation x = x `elem` ".?!-;\'\"\\|"
count xs = (head xs, length xs)
