# Link: https://leetcode-cn.com/problems/reformat-department-table

SELECT `id`,
       SUM(IF(`month` = 'Jan', `revenue`, NULL)) `Jan_Revenue`,
       SUM(IF(`month` = 'Feb', `revenue`, NULL)) `Feb_Revenue`,
       SUM(IF(`month` = 'Mar', `revenue`, NULL)) `Mar_Revenue`,
       SUM(IF(`month` = 'Apr', `revenue`, NULL)) `Apr_Revenue`,
       SUM(IF(`month` = 'May', `revenue`, NULL)) `May_Revenue`,
       SUM(IF(`month` = 'Jun', `revenue`, NULL)) `Jun_Revenue`,
       SUM(IF(`month` = 'Jul', `revenue`, NULL)) `Jul_Revenue`,
       SUM(IF(`month` = 'Aug', `revenue`, NULL)) `Aug_Revenue`,
       SUM(IF(`month` = 'Sep', `revenue`, NULL)) `Sep_Revenue`,
       SUM(IF(`month` = 'Oct', `revenue`, NULL)) `Oct_Revenue`,
       SUM(IF(`month` = 'Nov', `revenue`, NULL)) `Nov_Revenue`,
       SUM(IF(`month` = 'Dec', `revenue`, NULL)) `Dec_Revenue`
FROM `department`
GROUP BY `id`
ORDER BY `id`;