module Main where
import qualified Data.List as DL
import qualified Debug.Trace as DT
-- jelly tree.

padding n = concat $ take n (repeat " ")
pretty Nil d = (padding $ d) ++ "-\n"
pretty (Node g lb rb) d = 
    concat [ padding d, show g, "\n"    
           , padding d, pretty lb (d+1) 
           , padding d, pretty rb (d+1) ]

grp_min (One n) = n
grp_min (Two n _) = n
grp_min (San n _ _) = n

grp_max (One n) = n
grp_max (Two _ n) = n
grp_max (San _ _ n) = n

grp_rm_max (Two m n) = One m
grp_rm_max (One m) = EmptyGroup

grp_rm_min (Two m n) = One n
grp_rm_min (One m) = EmptyGroup

grp_insert :: (Ord a) => Group a -> a -> Group a
grp_insert EmptyGroup n = One n
grp_insert (One a) n = (Two (min a n) (max a n))
grp_insert (Two a b) n = San x y z
    where
      x = min n $ min a b
      z = max n $ max a b
      y = if a < b && b < n
          then b
          else if b < n && n < a
               then n
               else a
          
grp_insert g _ = GroupErr "Tried to grp insert a San"          

j_new n = Node (One n) Nil Nil
leaf (Node _ Nil Nil) = True
leaf _ = False

data Group a = EmptyGroup
             | One a
             | Two a a
             | San a a a
             | GroupErr String
               deriving (Show, Eq, Ord)

data Jtree a = Node (Group a) (Jtree a) (Jtree a)
             | Nil
               deriving (Show, Eq, Ord)

j_insert Nil n = j_new n
j_insert (Node g left right) n = 
    split $ Node (grp_insert g n) left right

split (Node (San a b c) lb rb) = 
    Node (One b) (j_insert lb a) (j_insert rb c)
split n = n

split_left (Node (Two a b) lb rb) = 
    Node (One b) (j_insert lb a) rb

split_right (Node (Two a b) lb rb) = 
    Node (One a) lb (j_insert rb b)

tree_min_element (Node g lb _) = if lb == Nil 
                                 then grp_min g
                                 else tree_min_element lb

tree_max_element (Node g lb rb) = if rb == Nil 
                                  then grp_max g 
                                  else tree_max_element rb

maxrm t@(Node g lb rb) = if leaf t
                         then Node (grp_rm_max g) Nil Nil
                         else Node g lb (maxrm rb)

minrm t@(Node g lb rb) = if leaf t
                         then Node (grp_rm_min g) Nil Nil
                         else Node g (minrm lb) rb

fix_left n@(Node g lb rb) = 
    if lb == Nil then n else 
        let gmin = grp_min g
            maxEl = tree_max_element lb
        in 
          -- check to make sure the min of the group is larger than 
          -- the max val in left branch.
          if gmin > maxEl
          then n          
          else 
              -- remove the max val from the left branch
              -- remove the min val of the group
              let newLb = maxrm lb
                  newGrp = grp_rm_min g
              in                    
                -- insert the left-branch-max-val into the group
                -- insert the group-min-val into the left branch
                Node (grp_insert newGrp maxEl) (j_insert newLb maxEl) rb

fix_right n@(Node g lb rb) = 
    if rb == Nil then n else 
        let gmax = grp_max g
            minEl = tree_min_element rb 
        in 
          -- check to make sure the max of the group is smaller than 
          -- the min val in right branch.
          if gmax < minEl
          then n          
          else 
              -- remove the min val from the right branch
              -- remove the max val of the group
              let newRb = minrm rb
                  newGrp = grp_rm_max g
              in                    
                -- insert the right-branch-min-val into the group
                -- insert the group-max-val into the right branch
                Node (grp_insert newGrp minEl) lb (j_insert newRb gmax)

tree_insert t el = j_insert (fix_left $ fix_right t) el
                   
main = do
  let tree = DL.foldl' tree_insert (j_new 0) [1..8000000]
  print $ tree_max_element tree
           
 
-- 100.000 0.523
-- 150.000 0.786
-- 200.000 1.071
-- 250.000 1.
-- 1.000.000 6.229
-- 2.000.000 11.992
-- 4.000.000 23.590
-- 8.000.000