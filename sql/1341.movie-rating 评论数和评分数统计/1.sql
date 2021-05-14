# Link: https://leetcode-cn.com/problems/movie-rating

(SELECT `name` `results`
 FROM `users`                 `u`
          JOIN `movie_rating` `mr`
               ON `u`.`user_id` = `mr`.`user_id`
 GROUP BY `u`.`user_id`
 ORDER BY COUNT(*) DESC, `name`
 LIMIT 1)
UNION
(SELECT `title` `results`
 FROM `movies`                `m`
          JOIN `movie_rating` `mr`
               ON `m`.`movie_id` = `mr`.`movie_id`
 WHERE LEFT(`created_at`, 7) = '2020-02'
 GROUP BY `m`.`movie_id`
 ORDER BY AVG(`rating`) DESC, `title`
 LIMIT 1);