# Link: https://leetcode.cn/problems/not-boring-movies

SELECT `id`, `movie`, `description`, `rating`
FROM `cinema`
WHERE `description` != 'boring'
  AND `id` % 2 = 1
ORDER BY `rating` DESC