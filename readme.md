# LeetCode

[![](https://img.shields.io/badge/Language-Go-%2300ADD8)](https://golang.org/)
[![](https://img.shields.io/badge/AC-673-%23F781BE)](https://leetcode-cn.com/u/bygo/)

---

# Template
- [Tree](https://github.com/bygo/leetcode/tree/master/templates/tree)
- [LinkedList](https://github.com/bygo/leetcode/tree/master/templates/linked_list)
- [Stack](https://github.com/bygo/leetcode/tree/master/templates/stack)
- [Array](https://github.com/bygo/leetcode/tree/master/templates/array)

---
- [TwoPointer](#TwoPointer)
- [Sql](#Sql)
- [Bit](#Bit)
- [Bfs](#Bfs)
- [Stack](#Stack)
- [Enum](#Enum)
- [Tree](#Tree)
- [Divide&conquer](#Divide&conquer)
- [Back](#Back)
- [Sort](#Sort)
- [Search](#Search)
- [LinkedList](#LinkedList)
- [BinarySearch](#BinarySearch)
- [Catalan](#Catalan)
- [StackMonotonic](#StackMonotonic)
- [SortCounter](#SortCounter)
- [Hash](#Hash)
- [Greedy](#Greedy)
- [Dp](#Dp)
- [Classic](#Classic)
- [Math](#Math)

---



## TwoPointer
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0003 | [Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) | 37.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0003.longest-substring-without-repeating-characters%20无重复字符的最长子串/main.go)| 
| 0019 | [Remove Nth Node From End of List](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list) | 42.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0019.remove-nth-node-from-end-of-list%20删除链表的倒数第%20N%20个结点/main.go)| 
| 0027 | [Remove Element                  ](https://leetcode-cn.com/problems/remove-element) | 59.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0027.remove-element%20移除元素/main.go)| 
| 0030 | [Substring with Concatenation of All Words](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words) | 35.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0030.substring-with-concatenation-of-all-words%20串联所有单词的子串/main.go)| 
| 0042 | [Trapping Rain Water             ](https://leetcode-cn.com/problems/trapping-rain-water) | 57.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0042.trapping-rain-water%20接雨水/main.go)| 
| 0075 | [Sort Colors                     ](https://leetcode-cn.com/problems/sort-colors) | 59.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0075.sort-colors%20颜色分类/main.go)| 
| 0080 | [Remove Duplicates from Sorted Array II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii) | 61.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0080.remove-duplicates-from-sorted-array-ii%20删除有序数组中的重复项%20II/main.go)| 
| 0443 | [String Compression              ](https://leetcode-cn.com/problems/string-compression) | 47.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0443.string-compression%20压缩字符串/main.go)| 
| 0930 | [Binary Subarrays With Sum       ](https://leetcode-cn.com/problems/binary-subarrays-with-sum) | 54.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0930.binary-subarrays-with-sum%20和相同的二元子数组/main.go)| 
| 1790 | [Check if One String Swap Can Make Strings Equal](https://leetcode-cn.com/problems/check-if-one-string-swap-can-make-strings-equal) | 65.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/1790.check-if-one-string-swap-can-make-strings-equal%20仅执行一次字符串交换能否使两个字符串相等/main.go)| 

## Sql
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0175 | [Combine Two Tables              ](https://leetcode-cn.com/problems/combine-two-tables) | 74.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0175.combine-two-tables%20给所有人加上地址/1.sql)| 
| 0176 | [Second Highest Salary           ](https://leetcode-cn.com/problems/second-highest-salary) | 35.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0176.second-highest-salary%20找第二名%20Null%20也返回/1.sql)| 
| 0177 | [Nth Highest Salary              ](https://leetcode-cn.com/problems/nth-highest-salary) | 46.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0177.nth-highest-salary%20找第N名%20Null%20也返回/1.sql)| 
| 0178 | [Rank Scores                     ](https://leetcode-cn.com/problems/rank-scores) | 60.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0178.rank-scores%20稠密排行/1.sql)| 
| 0180 | [Consecutive Numbers             ](https://leetcode-cn.com/problems/consecutive-numbers) | 48.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0180.consecutive-numbers%20连续出现3次的数字/1.sql)| 
| 0181 | [Employees Earning More Than Their Managers](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers) | 70.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0181.employees-earning-more-than-their-managers%20超过经理收入的员工/1.sql)| 
| 0182 | [Duplicate Emails                ](https://leetcode-cn.com/problems/duplicate-emails) | 79.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0182.duplicate-emails%20找重复邮箱/1.sql)| 
| 0183 | [Customers Who Never Order       ](https://leetcode-cn.com/problems/customers-who-never-order) | 67.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0183.customers-who-never-order%20从不订购的客户/1.sql)| 
| 0184 | [Department Highest Salary       ](https://leetcode-cn.com/problems/department-highest-salary) | 47.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0184.department-highest-salary%20部门最高工资的员工/1.sql)| 
| 0185 | [Department Top Three Salaries   ](https://leetcode-cn.com/problems/department-top-three-salaries) | 49.3% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0185.department-top-three-salaries%20部门工资前三高的员工/1.sql)| 
| 0196 | [Delete Duplicate Emails         ](https://leetcode-cn.com/problems/delete-duplicate-emails) | 66.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0196.delete-duplicate-emails%20删除重复邮箱/1.sql)| 
| 0197 | [Rising Temperature              ](https://leetcode-cn.com/problems/rising-temperature) | 53.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0197.rising-temperature%20温度相比昨天是上升的/1.sql)| 
| 0262 | [Trips and Users                 ](https://leetcode-cn.com/problems/trips-and-users) | 43.8% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0262.trips-and-users%20非禁止用户取消率/1.sql)| 
| 0511 | [Game Play Analysis I            ](https://leetcode-cn.com/problems/game-play-analysis-i) | 73.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0511.game-play-analysis-i%20首次登陆的时间/1.sql)| 
| 0512 | [Game Play Analysis II           ](https://leetcode-cn.com/problems/game-play-analysis-ii) | 53.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0512.game-play-analysis-ii%20首次登陆的设备名称/1.sql)| 
| 0534 | [Game Play Analysis III          ](https://leetcode-cn.com/problems/game-play-analysis-iii) | 67.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0534.game-play-analysis-iii%20每人每天累积多少时长/1.sql)| 
| 0550 | [Game Play Analysis IV           ](https://leetcode-cn.com/problems/game-play-analysis-iv) | 43.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0550.game-play-analysis-iv%20首日后隔天登录玩家的比率/1.sql)| 
| 0569 | [Median Employee Salary          ](https://leetcode-cn.com/problems/median-employee-salary) | 56.2% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0569.median-employee-salary%20每个公司薪酬中位数/1.sql)| 
| 0570 | [Managers with at Least 5 Direct Reports](https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports) | 66.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0570.managers-with-at-least-5-direct-reports%20至少5名下属的经理/1.sql)| 
| 0571 | [Find Median Given Frequency of Numbers](https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers) | 48.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0571.find-median-given-frequency-of-numbers%20频率数字的中位数/1.sql)| 
| 0574 | [Winning Candidate               ](https://leetcode-cn.com/problems/winning-candidate) | 64.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0574.winning-candidate%20当选者/1.sql)| 
| 0577 | [Employee Bonus                  ](https://leetcode-cn.com/problems/employee-bonus) | 70.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0577.employee-bonus%20员工奖金/1.sql)| 
| 0578 | [Get Highest Answer Rate Question](https://leetcode-cn.com/problems/get-highest-answer-rate-question) | 49.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0578.get-highest-answer-rate-question%20回答率最高的问题/1.sql)| 
| 0579 | [Find Cumulative Salary of an Employee](https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee) | 41.3% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0579.find-cumulative-salary-of-an-employee%20员工累积薪水/1.sql)| 
| 0580 | [Count Student Number in Departments](https://leetcode-cn.com/problems/count-student-number-in-departments) | 52.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0580.count-student-number-in-departments%20各专业学生人数/1.sql)| 
| 0584 | [Find Customer Referee           ](https://leetcode-cn.com/problems/find-customer-referee) | 77.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0584.find-customer-referee%20用户的推荐人/1.sql)| 
| 0585 | [Investments in 2016             ](https://leetcode-cn.com/problems/investments-in-2016) | 59.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0585.investments-in-2016%202016年的投资/1.sql)| 
| 0586 | [Customer Placing the Largest Number of Orders](https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders) | 75.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0586.customer-placing-the-largest-number-of-orders%20订单最多的客户/1.sql)| 
| 0595 | [Big Countries                   ](https://leetcode-cn.com/problems/big-countries) | 79.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0595.big-countries%20大的国家/1.sql)| 
| 0596 | [Classes More Than 5 Students    ](https://leetcode-cn.com/problems/classes-more-than-5-students) | 42.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0596.classes-more-than-5-students%20超过5名学生的课/1.sql)| 
| 0597 | [Friend Requests I: Overall Acceptance Rate](https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate) | 44.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0597.friend-requests-i-overall-acceptance-rate%20好友申请总体通过率/1.sql)| 
| 0601 | [Human Traffic of Stadium        ](https://leetcode-cn.com/problems/human-traffic-of-stadium) | 49.7% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0601.human-traffic-of-stadium%20人流量/1.sql)| 
| 0602 | [Friend Requests II: Who Has the Most Friends](https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends) | 62.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0602.friend-requests-ii-who-has-the-most-friends%20谁有最多的好友/1.sql)| 
| 0603 | [Consecutive Available Seats     ](https://leetcode-cn.com/problems/consecutive-available-seats) | 67.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0603.consecutive-available-seats%20连续空余座位/1.sql)| 
| 0607 | [Sales Person                    ](https://leetcode-cn.com/problems/sales-person) | 66.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0607.sales-person%20销售员/1.sql)| 
| 0608 | [Tree Node                       ](https://leetcode-cn.com/problems/tree-node) | 65.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0608.tree-node%20树节点/1.sql)| 
| 0610 | [Triangle Judgement              ](https://leetcode-cn.com/problems/triangle-judgement) | 66.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0610.triangle-judgement%20判断三角形/1.sql)| 
| 0612 | [Shortest Distance in a Plane    ](https://leetcode-cn.com/problems/shortest-distance-in-a-plane) | 66.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0612.shortest-distance-in-a-plane%20平面上的最近距离/1.sql)| 
| 0613 | [Shortest Distance in a Line     ](https://leetcode-cn.com/problems/shortest-distance-in-a-line) | 81.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0613.shortest-distance-in-a-line%20直线上的最近距离/1.sql)| 
| 0614 | [Second Degree Follower          ](https://leetcode-cn.com/problems/second-degree-follower) | 34.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0614.second-degree-follower%20二级关注者/1.sql)| 
| 0615 | [Average Salary: Departments VS Company](https://leetcode-cn.com/problems/average-salary-departments-vs-company) | 42.4% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0615.average-salary-departments-vs-company%20平均工资：部门与公司比较/1.sql)| 
| 0618 | [Students Report By Geography    ](https://leetcode-cn.com/problems/students-report-by-geography) | 62.1% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0618.students-report-by-geography%20学生地理信息报告/1.sql)| 
| 0619 | [Biggest Single Number           ](https://leetcode-cn.com/problems/biggest-single-number) | 46.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0619.biggest-single-number%20最出现一次的最大数字/1.sql)| 
| 0620 | [Not Boring Movies               ](https://leetcode-cn.com/problems/not-boring-movies) | 77.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0620.not-boring-movies%20有趣的电影/1.sql)| 
| 0626 | [Exchange Seats                  ](https://leetcode-cn.com/problems/exchange-seats) | 68.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0626.exchange-seats%20换座位/1.sql)| 
| 0627 | [Swap Salary                     ](https://leetcode-cn.com/problems/swap-salary) | 80.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0627.swap-salary%20变更性别/1.sql)| 
| 1045 | [Customers Who Bought All Products](https://leetcode-cn.com/problems/customers-who-bought-all-products) | 64.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1045.customers-who-bought-all-products%20买下所有产品的客户/1.sql)| 
| 1050 | [Actors and Directors Who Cooperated At Least Three Times](https://leetcode-cn.com/problems/actors-and-directors-who-cooperated-at-least-three-times) | 75.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1050.actors-and-directors-who-cooperated-at-least-three-times%20合作至少三次的演员和导员/1.sql)| 
| 1068 | [Product Sales Analysis I        ](https://leetcode-cn.com/problems/product-sales-analysis-i) | 86.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1068.product-sales-analysis-i%20产品的年份和价格/1.sql)| 
| 1069 | [Product Sales Analysis II       ](https://leetcode-cn.com/problems/product-sales-analysis-ii) | 82.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1069.product-sales-analysis-ii%20产品的销售总额/1.sql)| 
| 1070 | [Product Sales Analysis III      ](https://leetcode-cn.com/problems/product-sales-analysis-iii) | 48.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1070.product-sales-analysis-iii%20第一年的价格/1.sql)| 
| 1075 | [Project Employees I             ](https://leetcode-cn.com/problems/project-employees-i) | 69.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1075.project-employees-i%20项目的平均年限/1.sql)| 
| 1076 | [Project Employees II            ](https://leetcode-cn.com/problems/project-employees-ii) | 49.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1076.project-employees-ii%20员工最多的项目%20销售额最高的销售者/1.sql)| 
| 1077 | [Project Employees III           ](https://leetcode-cn.com/problems/project-employees-iii) | 72.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1077.project-employees-iii%20项目经济最丰富的员工/1.sql)| 
| 1082 | [Sales Analysis I                ](https://leetcode-cn.com/problems/sales-analysis-i) | 76.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1082.sales-analysis-i%20销售额最高的销售者/1.sql)| 
| 1083 | [Sales Analysis II               ](https://leetcode-cn.com/problems/sales-analysis-ii) | 52.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1083.sales-analysis-ii%20买%20S8%20却没有%20iPhone%20的买家/1.sql)| 
| 1084 | [Sales Analysis III              ](https://leetcode-cn.com/problems/sales-analysis-iii) | 53.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1084.sales-analysis-iii%20只在春季销售的产品/1.sql)| 
| 1097 | [Game Play Analysis V            ](https://leetcode-cn.com/problems/game-play-analysis-v) | 51.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1097.game-play-analysis-v%20第二天留存率/1.sql)| 
| 1098 | [Unpopular Books                 ](https://leetcode-cn.com/problems/unpopular-books) | 47.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1098.unpopular-books%20过去一年少于十本的书/1.sql)| 
| 1107 | [New Users Daily Count           ](https://leetcode-cn.com/problems/new-users-daily-count) | 41.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1107.new-users-daily-count%20每日新用户/1.sql)| 
| 1112 | [Highest Grade For Each Student  ](https://leetcode-cn.com/problems/highest-grade-for-each-student) | 64.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1112.highest-grade-for-each-student%20最高分数的学科/1.sql)| 
| 1113 | [Reported Posts                  ](https://leetcode-cn.com/problems/reported-posts) | 54.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1113.reported-posts%20不同的报告记录/1.sql)| 
| 1126 | [Active Businesses               ](https://leetcode-cn.com/problems/active-businesses) | 69.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1126.active-businesses%20活跃的业务/1.sql)| 
| 1127 | [User Purchase Platform          ](https://leetcode-cn.com/problems/user-purchase-platform) | 41.7% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1127.user-purchase-platform%20统计单端或双端人数/1.sql)| 
| 1132 | [Reported Posts II               ](https://leetcode-cn.com/problems/reported-posts-ii) | 38.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1132.reported-posts-ii%20垃圾清除率/1.sql)| 
| 1141 | [User Activity for the Past 30 Days I](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i) | 51.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1141.user-activity-for-the-past-30-days-i%20近30天活跃数/1.sql)| 
| 1142 | [User Activity for the Past 30 Days II](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii) | 37.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1142.user-activity-for-the-past-30-days-ii%20平均会话次数/1.sql)| 
| 1148 | [Article Views I                 ](https://leetcode-cn.com/problems/article-views-i) | 71.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1148.article-views-i%20浏览过自己文章的作者/1.sql)| 
| 1149 | [Article Views II                ](https://leetcode-cn.com/problems/article-views-ii) | 44.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1149.article-views-ii%20一天浏览过两篇文章的人/1.sql)| 
| 1158 | [Market Analysis I               ](https://leetcode-cn.com/problems/market-analysis-i) | 55.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1158.market-analysis-i%20统计2019年订单总数/1.sql)| 
| 1159 | [Market Analysis II              ](https://leetcode-cn.com/problems/market-analysis-ii) | 48.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1159.market-analysis-ii%20售出第二件不是喜欢的商品/1.sql)| 
| 1164 | [Product Price at a Given Date   ](https://leetcode-cn.com/problems/product-price-at-a-given-date) | 61.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1164.product-price-at-a-given-date%20变更后的价格/1.sql)| 
| 1173 | [Immediate Food Delivery I       ](https://leetcode-cn.com/problems/immediate-food-delivery-i) | 75.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1173.immediate-food-delivery-i%20当天配送率/1.sql)| 
| 1174 | [Immediate Food Delivery II      ](https://leetcode-cn.com/problems/immediate-food-delivery-ii) | 62.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1174.immediate-food-delivery-ii%20首日即时配送订单/1.sql)| 
| 1179 | [Reformat Department Table       ](https://leetcode-cn.com/problems/reformat-department-table) | 63.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1179.reformat-department-table%20格式化工资/1.sql)| 
| 1193 | [Monthly Transactions I          ](https://leetcode-cn.com/problems/monthly-transactions-i) | 61.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1193.monthly-transactions-i%20统计/1.sql)| 
| 1194 | [Tournament Winners              ](https://leetcode-cn.com/problems/tournament-winners) | 51.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1194.tournament-winners%20每组的优胜者/1.sql)| 
| 1204 | [Last Person to Fit in the Bus   ](https://leetcode-cn.com/problems/last-person-to-fit-in-the-bus) | 72.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1204.last-person-to-fit-in-the-elevator%20最后进入电梯/1.sql)| 
| 1205 | [Monthly Transactions II         ](https://leetcode-cn.com/problems/monthly-transactions-ii) | 46.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1205.monthly-transactions-ii%20交易统计/1.sql)| 
| 1211 | [Queries Quality and Percentage  ](https://leetcode-cn.com/problems/queries-quality-and-percentage) | 65.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1211.queries-quality-and-percentage/1.sql)| 
| 1212 | [Team Scores in Football Tournament](https://leetcode-cn.com/problems/team-scores-in-football-tournament) | 51.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1212.team-scores-in-football-tournament%20计算得分/1.sql)| 
| 1225 | [Report Contiguous Dates         ](https://leetcode-cn.com/problems/report-contiguous-dates) | 54.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1225.report-contiguous-dates%20统计连续/1.sql)| 
| 1241 | [Number of Comments per Post     ](https://leetcode-cn.com/problems/number-of-comments-per-post) | 59.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1241.number-of-comments-per-post%20查询评论数/1.sql)| 
| 1251 | [Average Selling Price           ](https://leetcode-cn.com/problems/average-selling-price) | 77.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1251.average-selling-price/1.sql)| 
| 1264 | [Page Recommendations            ](https://leetcode-cn.com/problems/page-recommendations) | 57.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1264.page-recommendations%20推荐朋友喜欢的页面/1.sql)| 
| 1270 | [All People Report to the Given Manager](https://leetcode-cn.com/problems/all-people-report-to-the-given-manager) | 80.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1270.all-people-report-to-the-given-manager%20递归查询/1.sql)| 
| 1280 | [Students and Examinations       ](https://leetcode-cn.com/problems/students-and-examinations) | 50.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1280.students-and-examinations%20学生各科测试次数/1.sql)| 
| 1285 | [Find the Start and End Number of Continuous Ranges](https://leetcode-cn.com/problems/find-the-start-and-end-number-of-continuous-ranges) | 81.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1285.find-the-start-and-end-number-of-continuous-ranges%20区间次数/1.sql)| 
| 1294 | [Weather Type in Each Country    ](https://leetcode-cn.com/problems/weather-type-in-each-country) | 66.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1294.weather-type-in-each-country%2011月份的天气/1.sql)| 
| 1303 | [Find the Team Size              ](https://leetcode-cn.com/problems/find-the-team-size) | 82.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1303.find-the-team-size%20团队人数/1.sql)| 
| 1308 | [Running Total for Different Genders](https://leetcode-cn.com/problems/running-total-for-different-genders) | 75.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1308.running-total-for-different-genders%20男女的累积分数/1.sql)| 
| 1321 | [Restaurant Growth               ](https://leetcode-cn.com/problems/restaurant-growth) | 61.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1321.restaurant-growth%20最近七天平均值/1.sql)| 
| 1322 | [Ads Performance                 ](https://leetcode-cn.com/problems/ads-performance) | 59.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1322.ads-performance%20广告效果统计/1.sql)| 
| 1327 | [List the Products Ordered in a Period](https://leetcode-cn.com/problems/list-the-products-ordered-in-a-period) | 69.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1327.list-the-products-ordered-in-a-period%20指定日期总数/1.sql)| 
| 1336 | [Number of Transactions per Visit](https://leetcode-cn.com/problems/number-of-transactions-per-visit) | 42.3% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1336.number-of-transactions-per-visit%20每次访问的交易次数/1.sql)| 
| 1341 | [Movie Rating                    ](https://leetcode-cn.com/problems/movie-rating) | 48.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1341.movie-rating%20评论数和评分数统计/1.sql)| 
| 1350 | [Students With Invalid Departments](https://leetcode-cn.com/problems/students-with-invalid-departments) | 84.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1350.students-with-invalid-departments%20不存在院校的学生/1.sql)| 
| 1355 | [Activity Participants           ](https://leetcode-cn.com/problems/activity-participants) | 66.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1355.activity-participants%20去除第一和倒数第一/1.sql)| 
| 1364 | [Number of Trusted Contacts of a Customer](https://leetcode-cn.com/problems/number-of-trusted-contacts-of-a-customer) | 67.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1364.number-of-trusted-contacts-of-a-customer%20互为好友的顾客/1.sql)| 
| 1369 | [Get the Second Most Recent Activity](https://leetcode-cn.com/problems/get-the-second-most-recent-activity) | 60.7% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1369.get-the-second-most-recent-activity%20倒数第二次活动/1.sql)| 
| 1378 | [Replace Employee ID With The Unique Identifier](https://leetcode-cn.com/problems/replace-employee-id-with-the-unique-identifier) | 86.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1378.replace-employee-id-with-the-unique-identifier%20使用uuid替换id/1.sql)| 
| 1384 | [Total Sales Amount by Year      ](https://leetcode-cn.com/problems/total-sales-amount-by-year) | 54.2% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1384.total-sales-amount-by-year%20年度统计/1.sql)| 
| 1393 | [Capital Gain/Loss               ](https://leetcode-cn.com/problems/capital-gainloss) | 85.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1393.capital-gainloss%20%20资本损益/1.sql)| 
| 1398 | [Customers Who Bought Products A and B but Not C](https://leetcode-cn.com/problems/customers-who-bought-products-a-and-b-but-not-c) | 77.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1398.customers-who-bought-products-a-and-b-but-not-c%20买AB不买C/1.sql)| 
| 1407 | [Top Travellers                  ](https://leetcode-cn.com/problems/top-travellers) | 66.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1407.top-travellers%20旅行距离排名/1.sql)| 
| 1412 | [Find the Quiet Students in All Exams](https://leetcode-cn.com/problems/find-the-quiet-students-in-all-exams) | 55.1% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1412.find-the-quiet-students-in-all-exams%20成绩中游学生/1.sql)| 
| 1418 | [Display Table of Food Orders in a Restaurant](https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant) | 73.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1418.display-table-of-food-orders-in-a-restaurant%20点菜展示表/main.go)| 
| 1421 | [NPV Queries                     ](https://leetcode-cn.com/problems/npv-queries) | 74.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1421.npv-queries%20连表/1.sql)| 
| 1435 | [Create a Session Bar Chart      ](https://leetcode-cn.com/problems/create-a-session-bar-chart) | 62.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1435.create-a-session-bar-chart%20分段汇总/1.sql)| 
| 1440 | [Evaluate Boolean Expression     ](https://leetcode-cn.com/problems/evaluate-boolean-expression) | 69.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1440.evaluate-boolean-expression%20连表比较/1.sql)| 
| 1445 | [Apples & Oranges                ](https://leetcode-cn.com/problems/apples-oranges) | 85.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1445.apples-oranges%20销量差/1.sql)| 
| 1454 | [Active Users                    ](https://leetcode-cn.com/problems/active-users) | 40.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1454.active-users%20连续7天在线的用户/1.sql)| 
| 1459 | [Rectangles Area                 ](https://leetcode-cn.com/problems/rectangles-area) | 59.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1459.rectangles-area%20组合成矩形/1.sql)| 
| 1468 | [Calculate Salaries              ](https://leetcode-cn.com/problems/calculate-salaries) | 71.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1468.calculate-salaries%20扣税/1.sql)| 
| 1479 | [Sales by Day of the Week        ](https://leetcode-cn.com/problems/sales-by-day-of-the-week) | 57.4% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1479.sales-by-day-of-the-week%20每周销量统计/1.sql)| 
| 1485 | [Group Sold Products By The Date ](https://leetcode-cn.com/problems/group-sold-products-by-the-date) | 66.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1485.group-sold-products-by-the-date%20group聚合/1.sql)| 
| 1495 | [Friendly Movies Streamed Last Month](https://leetcode-cn.com/problems/friendly-movies-streamed-last-month) | 54.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1495.friendly-movies-streamed-last-month%20儿童适宜的电影/1.sql)| 
| 1501 | [Countries You Can Safely Invest In](https://leetcode-cn.com/problems/countries-you-can-safely-invest-in) | 59.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1501.countries-you-can-safely-invest-in%20通话时长高于平均/1.sql)| 
| 1511 | [Customer Order Frequency        ](https://leetcode-cn.com/problems/customer-order-frequency) | 68.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1511.customer-order-frequency%20六七月消费大于等于100/1.sql)| 
| 1517 | [Find Users With Valid E-Mails   ](https://leetcode-cn.com/problems/find-users-with-valid-e-mails) | 72.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1517.find-users-with-valid-e-mails%20邮箱正则/1.sql)| 
| 1527 | [Patients With a Condition       ](https://leetcode-cn.com/problems/patients-with-a-condition) | 50.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1527.patients-with-a-condition%20患病患者/1.sql)| 
| 1532 | [The Most Recent Three Orders    ](https://leetcode-cn.com/problems/the-most-recent-three-orders) | 62.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1532.the-most-recent-three-orders%20最近三笔订单/1.sql)| 
| 1543 | [Fix Product Name Format         ](https://leetcode-cn.com/problems/fix-product-name-format) | 54.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1543.fix-product-name-format%20格式化/1.sql)| 
| 1549 | [The Most Recent Orders for Each Product](https://leetcode-cn.com/problems/the-most-recent-orders-for-each-product) | 66.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1549.the-most-recent-orders-for-each-product%20产品最新订单/1.sql)| 
| 1555 | [Bank Account Summary            ](https://leetcode-cn.com/problems/bank-account-summary) | 43.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1555.bank-account-summary%20是否透支/1.sql)| 
| 1565 | [Unique Orders and Customers Per Month](https://leetcode-cn.com/problems/unique-orders-and-customers-per-month) | 72.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1565.unique-orders-and-customers-per-month%20月订单数和顾客数/1.sql)| 
| 1571 | [Warehouse Manager               ](https://leetcode-cn.com/problems/warehouse-manager) | 75.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1571.warehouse-manager%20计算立方/1.sql)| 
| 1581 | [Customer Who Visited but Did Not Make Any Transactions](https://leetcode-cn.com/problems/customer-who-visited-but-did-not-make-any-transactions) | 81.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1581.customer-who-visited-but-did-not-make-any-transactions%20进店未交易/1.sql)| 
| 1587 | [Bank Account Summary II         ](https://leetcode-cn.com/problems/bank-account-summary-ii) | 82.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1587.bank-account-summary-ii%20余大于10000的用户/1.sql)| 
| 1596 | [The Most Frequently Ordered Products for Each Customer](https://leetcode-cn.com/problems/the-most-frequently-ordered-products-for-each-customer) | 76.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1596.the-most-frequently-ordered-products-for-each-customer%20经常订购的商品/1.sql)| 
| 1607 | [Sellers With No Sales           ](https://leetcode-cn.com/problems/sellers-with-no-sales) | 54.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1607.sellers-with-no-sales%20没有卖出产品的销售/1.sql)| 
| 1613 | [Find the Missing IDs            ](https://leetcode-cn.com/problems/find-the-missing-ids) | 72.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1613.find-the-missing-ids%20缺失的id/1.sql)| 
| 1623 | [All Valid Triplets That Can Represent a Country](https://leetcode-cn.com/problems/all-valid-triplets-that-can-represent-a-country) | 78.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1623.all-valid-triplets-that-can-represent-a-country%20三个学校组合不同的代表队/1.sql)| 
| 1633 | [Percentage of Users Attended a Contest](https://leetcode-cn.com/problems/percentage-of-users-attended-a-contest) | 65.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1633.percentage-of-users-attended-a-contest%20注册率/1.sql)| 
| 1635 | [Hopper Company Queries I        ](https://leetcode-cn.com/problems/hopper-company-queries-i) | 54.4% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1635.hopper-company-queries-i%20活跃数与行程数/1.sql)| 
| 1645 | [Hopper Company Queries II       ](https://leetcode-cn.com/problems/hopper-company-queries-ii) | 41.2% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1645.hopper-company-queries-ii%20工作率/1.sql)| 
| 1651 | [Hopper Company Queries III      ](https://leetcode-cn.com/problems/hopper-company-queries-iii) | 68.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1651.hopper-company-queries-iii%20平均距离和平均时间/1.sql)| 
| 1661 | [Average Time of Process per Machine](https://leetcode-cn.com/problems/average-time-of-process-per-machine) | 73.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1661.average-time-of-process-per-machine%20平均运行时间/1.sql)| 
| 1667 | [Fix Names in a Table            ](https://leetcode-cn.com/problems/fix-names-in-a-table) | 60.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1667.fix-names-in-a-table%20%20修复名字/1.sql)| 
| 1677 | [Product's Worth Over Invoices   ](https://leetcode-cn.com/problems/products-worth-over-invoices) | 40.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1677.products-worth-over-invoices%20统计/1.sql)| 
| 1683 | [Invalid Tweets                  ](https://leetcode-cn.com/problems/invalid-tweets) | 89.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1683.invalid-tweets%20无效推文/1.sql)| 
| 1693 | [Daily Leads and Partners        ](https://leetcode-cn.com/problems/daily-leads-and-partners) | 80.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1693.daily-leads-and-partners%20统计/1.sql)| 
| 1699 | [Number of Calls Between Two Persons](https://leetcode-cn.com/problems/number-of-calls-between-two-persons) | 79.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1699.number-of-calls-between-two-persons%20通话次数/1.sql)| 
| 1709 | [Biggest Window Between Visits   ](https://leetcode-cn.com/problems/biggest-window-between-visits) | 71.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1709.biggest-window-between-visits%20日期缺失/1.sql)| 
| 1715 | [Count Apples and Oranges        ](https://leetcode-cn.com/problems/count-apples-and-oranges) | 73.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1715.count-apples-and-oranges%20统计/1.sql)| 
| 1729 | [Find Followers Count            ](https://leetcode-cn.com/problems/find-followers-count) | 62.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1729.find-followers-count%20关注数/1.sql)| 
| 1731 | [The Number of Employees Which Report to Each Employee](https://leetcode-cn.com/problems/the-number-of-employees-which-report-to-each-employee) | 47.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1731.the-number-of-employees-which-report-to-each-employee%20每位经理有多少下属/1.sql)| 
| 1741 | [Find Total Time Spent by Each Employee](https://leetcode-cn.com/problems/find-total-time-spent-by-each-employee) | 84.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1741.find-total-time-spent-by-each-employee%20员工花费时间/1.sql)| 
| 1747 | [Leetflex Banned Accounts        ](https://leetcode-cn.com/problems/leetflex-banned-accounts) | 70.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1747.leetflex-banned-accounts%20应被禁止的账户/1.sql)| 
| 1757 | [Recyclable and Low Fat Products ](https://leetcode-cn.com/problems/recyclable-and-low-fat-products) | 92.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1757.recyclable-and-low-fat-products%20可回收且低脂的产品/1.sql)| 
| 1767 | [Find the Subtasks That Did Not Execute](https://leetcode-cn.com/problems/find-the-subtasks-that-did-not-execute) | 79.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1767.find-the-subtasks-that-did-not-execute%20找到未执行的子任务/1.sql)| 
| 1777 | [Product's Price for Each Store  ](https://leetcode-cn.com/problems/products-price-for-each-store) | 77.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1777.products-price-for-each-store%20每家商店价格/1.sql)| 
| 1783 | [Grand Slam Titles               ](https://leetcode-cn.com/problems/grand-slam-titles) | 78.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1783.grand-slam-titles%20大满贯数/1.sql)| 
| 1789 | [Primary Department for Each Employee](https://leetcode-cn.com/problems/primary-department-for-each-employee) | 71.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1789.primary-department-for-each-employee%20员工的主要部门/1.sql)| 
| 1795 | [Rearrange Products Table        ](https://leetcode-cn.com/problems/rearrange-products-table) | 79.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1795.rearrange-products-table%20商品价格合并/1.sql)| 
| 1809 | [Ad-Free Sessions                ](https://leetcode-cn.com/problems/ad-free-sessions) | 64.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1809.ad-free-sessions%20没有展示广告的会话/1.sql)| 
| 1811 | [Find Interview Candidates       ](https://leetcode-cn.com/problems/find-interview-candidates) | 66.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1811.find-interview-candidates%20统计/1.sql)| 
| 1821 | [Find Customers With Positive Revenue this Year](https://leetcode-cn.com/problems/find-customers-with-positive-revenue-this-year) | 90.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1821.find-customers-with-positive-revenue-this-year%20正收入/1.sql)| 
| 1831 | [Maximum Transaction Each Day    ](https://leetcode-cn.com/problems/maximum-transaction-each-day) | 73.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1831.maximum-transaction-each-day%20每天最大交易/1.sql)| 

## Bit
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0168 | [Excel Sheet Column Title        ](https://leetcode-cn.com/problems/excel-sheet-column-title) | 43.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0168.excel-sheet-column-title%20Excel表列名称/main.go)| 
| 0231 | [Power of Two                    ](https://leetcode-cn.com/problems/power-of-two) | 50.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0231.power-of-two%202的幂/main.go)| 
| 0342 | [Power of Four                   ](https://leetcode-cn.com/problems/power-of-four) | 51.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0342.power-of-four%204的幂/main.go)| 
| 0401 | [Binary Watch                    ](https://leetcode-cn.com/problems/binary-watch) | 61.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0401.binary-watch%20二进制手表/main.go)| 
| 0461 | [Hamming Distance                ](https://leetcode-cn.com/problems/hamming-distance) | 81.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0461.hamming-distance%20汉明距离/main.go)| 
| 0477 | [Total Hamming Distance          ](https://leetcode-cn.com/problems/total-hamming-distance) | 60.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/0477.total-hamming-distance%20汉明距离总和/main.go)| 
| 1689 | [Partitioning Into Minimum Number Of Deci-Binary Numbers](https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers) | 87.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/1689.partitioning-into-minimum-number-of-deci-binary-numbers%20十-二进制数的最少数目/main.go)| 
| 0000 | [Offer 0015                      ](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof 二进制中1的个数) | NaN% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/bit/Offer%200015.er-jin-zhi-zhong-1de-ge-shu-lcof%20二进制中1的个数/main.go)| 

## Bfs
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0098 | [Validate Binary Search Tree     ](https://leetcode-cn.com/problems/validate-binary-search-tree) | 34.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0098.validate-binary-search-tree%20验证二叉搜索树/main.go)| 
| 0102 | [Binary Tree Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) | 64.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0102.binary-tree-level-order-traversal%20二叉树的层序遍历/main.go)| 
| 0103 | [Binary Tree Zigzag Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) | 57.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0103.binary-tree-zigzag-level-order-traversal%20二叉树的锯齿形层序遍历/main.go)| 
| 0107 | [Binary Tree Level Order Traversal II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) | 69.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0107.binary-tree-level-order-traversal-ii%20二叉树的层序遍历%20II/main.go)| 
| 0111 | [Minimum Depth of Binary Tree    ](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) | 48.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0111.minimum-depth-of-binary-tree%20二叉树的最小深度/main.go)| 
| 0116 | [Populating Next Right Pointers in Each Node](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node) | 70.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0116.populating-next-right-pointers-in-each-node%20填充每个节点的下一个右侧节点指针/main.go)| 
| 0117 | [Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii) | 61.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0117.populating-next-right-pointers-in-each-node-ii%20填充每个节点的下一个右侧节点指针%20II/main.go)| 
| 0199 | [Binary Tree Right Side View     ](https://leetcode-cn.com/problems/binary-tree-right-side-view) | 65.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0199.binary-tree-right-side-view%20二叉树的右视图/main.go)| 
| 0297 | [Serialize and Deserialize Binary Tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) | 55.8% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0297.serialize-and-deserialize-binary-tree%20二叉树的序列化与反序列化/main.go)| 
| 0428 | [Serialize and Deserialize N-ary Tree](https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree) | 65.0% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0428.serialize-and-deserialize-n-ary-tree%20序列化和反序列化%20N%20叉树/main.go)| 
| 0429 | [N-ary Tree Level Order Traversal](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal) | 69.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0429.n-ary-tree-level-order-traversal%20N%20叉树的层序遍历/main.go)| 
| 0513 | [Find Bottom Left Tree Value     ](https://leetcode-cn.com/problems/find-bottom-left-tree-value) | 73.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0513.find-bottom-left-tree-value%20找树左下角的值/main.go)| 
| 0515 | [Find Largest Value in Each Tree Row](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row) | 64.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0515.find-largest-value-in-each-tree-row%20在每个树行中找最大值/main.go)| 
| 0559 | [Maximum Depth of N-ary Tree     ](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree) | 72.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0559.maximum-depth-of-n-ary-tree%20N%20叉树的最大深度/main.go)| 
| 0623 | [Add One Row to Tree             ](https://leetcode-cn.com/problems/add-one-row-to-tree) | 54.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0623.add-one-row-to-tree%20在二叉树中增加一行/main.go)| 
| 0637 | [Average of Levels in Binary Tree](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree) | 68.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0637.average-of-levels-in-binary-tree%20二叉树的层平均值/main.go)| 
| 1161 | [Maximum Level Sum of a Binary Tree](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree) | 63.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/1161.maximum-level-sum-of-a-binary-tree%20最大层内元素和/main.go)| 
| 1306 | [Jump Game III                   ](https://leetcode-cn.com/problems/jump-game-iii) | 57.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/1306.jump-game-iii%20跳跃游戏%20III/main.go)| 

## Stack
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0020 | [Valid Parentheses               ](https://leetcode-cn.com/problems/valid-parentheses) | 44.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0020.valid-parentheses%20有效(){}[]括号/main.go)| 
| 0032 | [Longest Valid Parentheses       ](https://leetcode-cn.com/problems/longest-valid-parentheses) | 35.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0032.longest-valid-parentheses%20最长有效()括号/main.go)| 
| 0071 | [Simplify Path                   ](https://leetcode-cn.com/problems/simplify-path) | 42.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0071.simplify-path%20简化路径/main.go)| 
| 0084 | [Largest Rectangle in Histogram  ](https://leetcode-cn.com/problems/largest-rectangle-in-histogram) | 43.3% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0084.largest-rectangle-in-histogram%20统计矩形/main.go)| 
| 0085 | [Maximal Rectangle               ](https://leetcode-cn.com/problems/maximal-rectangle) | 51.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0085.maximal-rectangle%20最大矩形/main.go)| 
| 0094 | [Binary Tree Inorder Traversal   ](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) | 75.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0094.binary-tree-inorder-traversal%20二叉树的中序遍历/main.go)| 
| 0103 | [Binary Tree Zigzag Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) | 57.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0103.binary-tree-zigzag-level-order-traversal%20二叉树的锯齿形层序遍历/main.go)| 
| 0144 | [Binary Tree Preorder Traversal  ](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) | 70.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0144.binary-tree-preorder-traversal%20二叉树的前序遍历/main.go)| 
| 0145 | [Binary Tree Postorder Traversal ](https://leetcode-cn.com/problems/binary-tree-postorder-traversal) | 74.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0145.binary-tree-postorder-traversal%20二叉树的后序遍历/main.go)| 
| 0150 | [Evaluate Reverse Polish Notation](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation) | 53.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0150.evaluate-reverse-polish-notation%20逆波兰表达式求值/main.go)| 
| 0155 | [Min Stack                       ](https://leetcode-cn.com/problems/min-stack) | 57.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0155.min-stack%20最小栈/main.go)| 
| 0224 | [Basic Calculator                ](https://leetcode-cn.com/problems/basic-calculator) | 41.8% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0224.basic-calculator%20加减小括号/main.go)| 
| 0225 | [Implement Stack using Queues    ](https://leetcode-cn.com/problems/implement-stack-using-queues) | 67.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0225.implement-stack-using-queues%20用队列实现栈/main.go)| 
| 0227 | [Basic Calculator II             ](https://leetcode-cn.com/problems/basic-calculator-ii) | 43.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0227.basic-calculator-ii%20加减乘除/main.go)| 
| 0232 | [Implement Queue using Stacks    ](https://leetcode-cn.com/problems/implement-queue-using-stacks) | 69.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0232.implement-queue-using-stacks%20用栈实现队列/main.go)| 
| 0255 | [Verify Preorder Sequence in Binary Search Tree](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree) | 47.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0255.verify-preorder-sequence-in-binary-search-tree%20验证前序遍历序列二叉搜索树/main.go)| 
| 0272 | [Closest Binary Search Tree Value II](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii) | 64.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0272.closest-binary-search-tree-value-ii%20最接近的二叉搜索树值%20II/main.go)| 
| 0316 | [Remove Duplicate Letters        ](https://leetcode-cn.com/problems/remove-duplicate-letters) | 47.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0316.remove-duplicate-letters%20去除重复字母/main.go)| 
| 0331 | [Verify Preorder Serialization of a Binary Tree](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree) | 48.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0331.verify-preorder-serialization-of-a-binary-tree%20验证二叉树的前序序列化/main.go)| 
| 0341 | [Flatten Nested List Iterator    ](https://leetcode-cn.com/problems/flatten-nested-list-iterator) | 72.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0341.flatten-nested-list-iterator%20扁平化嵌套列表迭代器/main.go)| 
| 0385 | [Mini Parser                     ](https://leetcode-cn.com/problems/mini-parser) | 41.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0385.mini-parser%20迷你语法分析器/main.go)| 
| 0394 | [Decode String                   ](https://leetcode-cn.com/problems/decode-string) | 55.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0394.decode-string%20字符串解码/main.go)| 
| 0402 | [Remove K Digits                 ](https://leetcode-cn.com/problems/remove-k-digits) | 32.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0402.remove-k-digits%20移掉K位数字后最小/main.go)| 
| 0430 | [Flatten a Multilevel Doubly Linked List](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list) | 54.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0430.flatten-a-multilevel-doubly-linked-list%20扁平化多级双向链表/main.go)| 
| 0439 | [Ternary Expression Parser       ](https://leetcode-cn.com/problems/ternary-expression-parser) | 58.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0439.ternary-expression-parser%20三元表达式解析器/main.go)| 
| 0445 | [Add Two Numbers II              ](https://leetcode-cn.com/problems/add-two-numbers-ii) | 58.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0445.add-two-numbers-ii%20两数相加/main.go)| 
| 0456 | [132 Pattern                     ](https://leetcode-cn.com/problems/132-pattern) | 36.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0456.132-pattern%20132模式/main.go)| 
| 0496 | [Next Greater Element I          ](https://leetcode-cn.com/problems/next-greater-element-i) | 68.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0496.next-greater-element-i%20下一个更大元素%20I/main.go)| 
| 0503 | [Next Greater Element II         ](https://leetcode-cn.com/problems/next-greater-element-ii) | 63.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0503.next-greater-element-ii%20下一个更大的值/main.go)| 
| 0591 | [Tag Validator                   ](https://leetcode-cn.com/problems/tag-validator) | 34.0% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0591.tag-validator%20标签验证器/main.go)| 
| 0636 | [Exclusive Time of Functions     ](https://leetcode-cn.com/problems/exclusive-time-of-functions) | 54.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0636.exclusive-time-of-functions%20函数的独占时间/main.go)| 
| 0726 | [Number of Atoms                 ](https://leetcode-cn.com/problems/number-of-atoms) | 55.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0726.number-of-atoms%20原子的数量/main.go)| 
| 0735 | [Asteroid Collision              ](https://leetcode-cn.com/problems/asteroid-collision) | 41.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0735.asteroid-collision%20星球碰撞/main.go)| 
| 0739 | [Daily Temperatures              ](https://leetcode-cn.com/problems/daily-temperatures) | 67.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0739.daily-temperatures%20能观测到更高温度的天数/main.go)| 
| 0772 | [Basic Calculator III            ](https://leetcode-cn.com/problems/basic-calculator-iii) | 49.0% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0772.basic-calculator-iii%20加减乘除小括号/main.go)| 
| 0844 | [Backspace String Compare        ](https://leetcode-cn.com/problems/backspace-string-compare) | 51.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0844.backspace-string-compare%20比较含退格的字符串/main.go)| 
| 0856 | [Score of Parentheses            ](https://leetcode-cn.com/problems/score-of-parentheses) | 62.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0856.score-of-parentheses%20括号的分数/main.go)| 
| 0880 | [Decoded String at Index         ](https://leetcode-cn.com/problems/decoded-string-at-index) | 25.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0880.decoded-string-at-index%20索引处的解码字符串/main.go)| 
| 0895 | [Maximum Frequency Stack         ](https://leetcode-cn.com/problems/maximum-frequency-stack) | 54.1% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0895.maximum-frequency-stack%20最大频率栈/main.go)| 
| 0901 | [Online Stock Span               ](https://leetcode-cn.com/problems/online-stock-span) | 54.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0901.online-stock-span%20股票价格跨度/main.go)| 
| 0921 | [Minimum Add to Make Parentheses Valid](https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid) | 73.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0921.minimum-add-to-make-parentheses-valid%20使括号有效的最少添加/main.go)| 
| 0946 | [Validate Stack Sequences        ](https://leetcode-cn.com/problems/validate-stack-sequences) | 62.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0946.validate-stack-sequences%20验证栈序列/main.go)| 
| 1003 | [Check If Word Is Valid After Substitutions](https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions) | 57.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1003.check-if-word-is-valid-after-substitutions%20检查替换后的词是否有效/main.go)| 
| 1019 | [Next Greater Node In Linked List](https://leetcode-cn.com/problems/next-greater-node-in-linked-list) | 59.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1019.next-greater-node-in-linked-list%20下一个更大的值/main.go)| 
| 1081 | [Smallest Subsequence of Distinct Characters](https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters) | 57.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1081.smallest-subsequence-of-distinct-characters%20去除重复字母/main.go)| 
| 1190 | [Reverse Substrings Between Each Pair of Parentheses](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses) | 65.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1190.reverse-substrings-between-each-pair-of-parentheses%20反转每对括号间的子串/main.go)| 
| 1544 | [Make The String Great           ](https://leetcode-cn.com/problems/make-the-string-great) | 54.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/1544.make-the-string-great%20整理字符串/main.go)| 
| 1598 | [Crawler Log Folder              ](https://leetcode-cn.com/problems/crawler-log-folder) | 67.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/1598.crawler-log-folder%20文件夹操作日志搜集器/main.go)| 

## Enum
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0002 | [Add Two Numbers                 ](https://leetcode-cn.com/problems/add-two-numbers) | 40.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0002.add-two-numbers%20两数相加/main.go)| 
| 0006 | [ZigZag Conversion               ](https://leetcode-cn.com/problems/zigzag-conversion) | 50.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0006.zigzag-conversion%20Z%20字形变换/main.go)| 
| 0014 | [Longest Common Prefix           ](https://leetcode-cn.com/problems/longest-common-prefix) | 40.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0014.longest-common-prefix%20最长公共前缀/main.go)| 
| 0036 | [Valid Sudoku                    ](https://leetcode-cn.com/problems/valid-sudoku) | 62.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0036.valid-sudoku%20有效的数独/main.go)| 
| 0038 | [Count and Say                   ](https://leetcode-cn.com/problems/count-and-say) | 58.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0038.count-and-say%20外观数列/main.go)| 
| 0041 | [First Missing Positive          ](https://leetcode-cn.com/problems/first-missing-positive) | 41.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/enum/0041.first-missing-positive%20缺失的第一个正数/main.go)| 
| 0121 | [Best Time to Buy and Sell Stock ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock) | 57.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0121.best-time-to-buy-and-sell-stock%20买卖股票的最佳时机/main.go)| 
| 0122 | [Best Time to Buy and Sell Stock II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii) | 68.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0122.best-time-to-buy-and-sell-stock-ii%20买卖股票的最佳时机%20II/main.go)| 
| 0171 | [Excel Sheet Column Number       ](https://leetcode-cn.com/problems/excel-sheet-column-number) | 71.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0171.excel-sheet-column-number%20Excel%20表列序号/main.go)| 
| 1213 | [Intersection of Three Sorted Arrays](https://leetcode-cn.com/problems/intersection-of-three-sorted-arrays) | 77.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1213.intersection-of-three-sorted-arrays%20三个有序数组的交集/main.go)| 
| 1295 | [Find Numbers with Even Number of Digits](https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits) | 80.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1295.find-numbers-with-even-number-of-digits%20统计位数为偶数的数字/main.go)| 
| 1389 | [Create Target Array in the Given Order](https://leetcode-cn.com/problems/create-target-array-in-the-given-order) | 83.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1389.create-target-array-in-the-given-order%20按既定顺序创建目标数组/main.go)| 
| 1503 | [Last Moment Before All Ants Fall Out of a Plank](https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank) | 51.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/1503.last-moment-before-all-ants-fall-out-of-a-plank%20所有蚂蚁掉下来前的最后一刻/main.go)| 
| 1646 | [Get Maximum in Generated Array  ](https://leetcode-cn.com/problems/get-maximum-in-generated-array) | 53.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1646.get-maximum-in-generated-array%20获取生成数组中的最大值/main.go)| 
| 1672 | [Richest Customer Wealth         ](https://leetcode-cn.com/problems/richest-customer-wealth) | 85.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1672.richest-customer-wealth%20最富有客户的资产总量/main.go)| 

## Tree
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0022 | [Generate Parentheses            ](https://leetcode-cn.com/problems/generate-parentheses) | 77.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0022.generate-parentheses/dfs/catalan/main.go)| dfs.catalan
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0022.generate-parentheses/dfs/recursive/main.go)| dfs.recursive
| 0039 | [Combination Sum                 ](https://leetcode-cn.com/problems/combination-sum) | 72.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0039.combination-sum/dfs/recursive/main.go)| dfs.recursive
| 0040 | [Combination Sum II              ](https://leetcode-cn.com/problems/combination-sum-ii) | 62.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0040.combination-sum-ii/dfs/recursive/main.go)| dfs.recursive
| 0094 | [Binary Tree Inorder Traversal   ](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) | 75.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/break/main.go)| dfs.morris.break
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/keep/main.go)| dfs.morris.keep
| 0095 | [Unique Binary Search Trees II   ](https://leetcode-cn.com/problems/unique-binary-search-trees-ii) | 69.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0095.unique-binary-search-trees-ii/dfs/catalan/main.go)| dfs.catalan
| 0096 | [Unique Binary Search Trees      ](https://leetcode-cn.com/problems/unique-binary-search-trees) | 69.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0096.unique-binary-search-trees/catalan/main.go)| catalan
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0096.unique-binary-search-trees/dp/main.go)| dp
| 0098 | [Validate Binary Search Tree     ](https://leetcode-cn.com/problems/validate-binary-search-tree) | 34.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/preorder/recursive/main.go)| dfs.preorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/preorder/stack/main.go)| dfs.preorder.stack
| 0099 | [Recover Binary Search Tree      ](https://leetcode-cn.com/problems/recover-binary-search-tree) | 61.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0099.recover-binary-search-tree/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0099.recover-binary-search-tree/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0099.recover-binary-search-tree/dfs/morris/keep.go)| dfs.morris
| 0100 | [Same Tree                       ](https://leetcode-cn.com/problems/same-tree) | 60.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0100.same-tree/dfs/recursive/main.go)| dfs.recursive
| 0101 | [Symmetric Tree                  ](https://leetcode-cn.com/problems/symmetric-tree) | 56.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0101.symmetric-tree/dfs/recursive/main.go)| dfs.recursive
| 0102 | [Binary Tree Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) | 64.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0102.binary-tree-level-order-traversal%20二叉树的层序遍历/dfs/recursive/main.go)| dfs.recursive
| 0103 | [Binary Tree Zigzag Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) | 57.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0103.binary-tree-zigzag-level-order-traversal%20二叉树的锯齿形层序遍历/dfs/recursive/main.go)| dfs.recursive
| 0104 | [Maximum Depth of Binary Tree    ](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) | 76.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0104.maximum-depth-of-binary-tree/dfs/main.go)| dfs
| 0105 | [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal) | 70.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0105.construct-binary-tree-from-preorder-and-inorder-traversal/main.go)| 
| 0106 | [Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal) | 72.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0106.construct-binary-tree-from-inorder-and-postorder-traversal/main.go)| 
| 0107 | [Binary Tree Level Order Traversal II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) | 69.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0107.binary-tree-level-order-traversal-ii/dfs/recursive/main.go)| dfs.recursive
| 0108 | [Convert Sorted Array to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree) | 75.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0108.convert-sorted-array-to-binary-search-tree/dfs/recursive/main.go)| dfs.recursive
| 0109 | [Convert Sorted List to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree) | 76.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0109.convert-sorted-list-to-binary-search-tree/array/main.go)| array
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0109.convert-sorted-list-to-binary-search-tree/inorder/main.go)| inorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0109.convert-sorted-list-to-binary-search-tree/recursive/main.go)| recursive
| 0110 | [Balanced Binary Tree            ](https://leetcode-cn.com/problems/balanced-binary-tree) | 56.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0110.balanced-binary-tree/postorder/main.go)| postorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0110.balanced-binary-tree/top/main.go)| top
| 0111 | [Minimum Depth of Binary Tree    ](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) | 48.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0111.minimum-depth-of-binary-tree/dfs/main.go)| dfs
| 0112 | [Path Sum                        ](https://leetcode-cn.com/problems/path-sum) | 52.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0112.path-sum/dfs/main.go)| dfs
| 0113 | [Path Sum II                     ](https://leetcode-cn.com/problems/path-sum-ii) | 62.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0113.path-sum-ii/dfs/main.go)| dfs
| 0114 | [Flatten Binary Tree to Linked List](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list) | 72.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0114.flatten-binary-tree-to-linked-list/preorder.morris/main.go)| preorder.morris
| 0116 | [Populating Next Right Pointers in Each Node](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node) | 70.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0116.populating-next-right-pointers-in-each-node/dfs/main.go)| dfs
| 0124 | [Binary Tree Maximum Path Sum    ](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum) | 44.3% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0124.binary-tree-maximum-path-sum/dfs/main.go)| dfs
| 0129 | [Sum Root to Leaf Numbers        ](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers) | 68.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0129.sum-root-to-leaf-numbers/dfs/main.go)| dfs
| 0144 | [Binary Tree Preorder Traversal  ](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) | 70.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/dfs/stack/main.go)| dfs.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/morris/break/main.go)| morris.break
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/morris/keep/main.go)| morris.keep
| 0145 | [Binary Tree Postorder Traversal ](https://leetcode-cn.com/problems/binary-tree-postorder-traversal) | 74.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0145.binary-tree-postorder-traversal/dfs/recursive/main.go)| dfs.recursive
| 0156 | [Binary Tree Upside Down         ](https://leetcode-cn.com/problems/binary-tree-upside-down) | 72.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0156.binary-tree-upside-down/link/main.go)| link
| 0173 | [Binary Search Tree Iterator     ](https://leetcode-cn.com/problems/binary-search-tree-iterator) | 80.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0173.binary-search-tree-iterator/main.go)| 
| 0199 | [Binary Tree Right Side View     ](https://leetcode-cn.com/problems/binary-tree-right-side-view) | 65.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0199.binary-tree-right-side-view/dfs/main.go)| dfs
| 0222 | [Count Complete Tree Nodes       ](https://leetcode-cn.com/problems/count-complete-tree-nodes) | 78.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0222.count-complete-tree-nodes/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0222.count-complete-tree-nodes/two_pointer/main.go)| two pointer
| 0226 | [Invert Binary Tree              ](https://leetcode-cn.com/problems/invert-binary-tree) | 78.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0226.invert-binary-tree/dfs/main.go)| dfs
| 0230 | [Kth Smallest Element in a BST   ](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst) | 74.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0230.kth-smallest-element-in-a-bst/dfs/main.go)| dfs
| 0235 | [Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree) | 66.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0235.lowest-common-ancestor-of-a-binary-search-tree/dfs/main.go)| dfs
| 0236 | [Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree) | 68.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0236.lowest-common-ancestor-of-a-binary-tree/dfs/main.go)| dfs
| 0250 | [Count Univalue Subtrees         ](https://leetcode-cn.com/problems/count-univalue-subtrees) | 63.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0250.count-univalue-subtrees/dfs/main.go)| dfs
| 0255 | [Verify Preorder Sequence in Binary Search Tree](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree) | 47.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0255.verify-preorder-sequence-in-binary-search-tree/main.go)| 
| 0257 | [Binary Tree Paths               ](https://leetcode-cn.com/problems/binary-tree-paths) | 68.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0257.binary-tree-paths/dfs/main.go)| dfs
| 0270 | [Closest Binary Search Tree Value](https://leetcode-cn.com/problems/closest-binary-search-tree-value) | 55.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0270.closest-binary-search-tree-value/main.go)| 
| 0272 | [Closest Binary Search Tree Value II](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii) | 64.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0272.closest-binary-search-tree-value-ii/main.go)| 
| 0285 | [Inorder Successor in BST        ](https://leetcode-cn.com/problems/inorder-successor-in-bst) | 63.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0285.inorder-successor-in-bst/main.go)| 
| 0297 | [Serialize and Deserialize Binary Tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) | 55.8% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0297.serialize-and-deserialize-binary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0297.serialize-and-deserialize-binary-tree/dfs/main.go)| dfs
| 0298 | [Binary Tree Longest Consecutive Sequence](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence) | 57.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0298.binary-tree-longest-consecutive-sequence/main.go)| 
| 0366 | [Find Leaves of Binary Tree      ](https://leetcode-cn.com/problems/find-leaves-of-binary-tree) | 76.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0366.find-leaves-of-binary-tree/main.go)| 
| 0404 | [Sum of Left Leaves              ](https://leetcode-cn.com/problems/sum-of-left-leaves) | 58.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0404.sum-of-left-leaves/dfs/main.go)| dfs
| 0426 | [Convert Binary Search Tree to Sorted Doubly Linked List](https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list) | 66.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0426.convert-binary-search-tree-to-sorted-doubly-linked-list/inorder/main.go)| inorder
| 0428 | [Serialize and Deserialize N-ary Tree](https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree) | 65.0% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0428.serialize-and-deserialize-n-ary-tree/dfs/main.go)| dfs
| 0437 | [Path Sum III                    ](https://leetcode-cn.com/problems/path-sum-iii) | 56.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0437.path-sum-iii/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0437.path-sum-iii/hash_table/main.go)| hash table
| 0449 | [Serialize and Deserialize BST   ](https://leetcode-cn.com/problems/serialize-and-deserialize-bst) | 55.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0449.serialize-and-deserialize-bst/dfs/preorder/main.go)| dfs.preorder
| 0450 | [Delete Node in a BST            ](https://leetcode-cn.com/problems/delete-node-in-a-bst) | 48.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0450.delete-node-in-a-bst/dfs/main.go)| dfs
| 0501 | [Find Mode in Binary Search Tree ](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree) | 51.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0501.find-mode-in-binary-search-tree/dfs/recursive/main.go)| dfs.recursive
| 0508 | [Most Frequent Subtree Sum       ](https://leetcode-cn.com/problems/most-frequent-subtree-sum) | 67.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0508.most-frequent-subtree-sum/dfs/main.go)| dfs
| 0510 | [Inorder Successor in BST II     ](https://leetcode-cn.com/problems/inorder-successor-in-bst-ii) | 60.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0510.inorder-successor-in-bst-ii/main.go)| 
| 0530 | [Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst) | 62.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0530.minimum-absolute-difference-in-bst/dfs/main.go)| dfs
| 0538 | [Convert BST to Greater Tree     ](https://leetcode-cn.com/problems/convert-bst-to-greater-tree) | 69.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0538.convert-bst-to-greater-tree/dfs/main.go)| dfs
| 0543 | [Diameter of Binary Tree         ](https://leetcode-cn.com/problems/diameter-of-binary-tree) | 54.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0543.diameter-of-binary-tree/dfs/main.go)| dfs
| 0559 | [Maximum Depth of N-ary Tree     ](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree) | 72.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0559.maximum-depth-of-n-ary-tree/dfs/main.go)| dfs
| 0563 | [Binary Tree Tilt                ](https://leetcode-cn.com/problems/binary-tree-tilt) | 60.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0563.binary-tree-tilt/dfs/main.go)| dfs
| 0572 | [Subtree of Another Tree         ](https://leetcode-cn.com/problems/subtree-of-another-tree) | 47.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0572.subtree-of-another-tree/dfs/main.go)| dfs
| 0589 | [N-ary Tree Preorder Traversal   ](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal) | 74.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0589.n-ary-tree-preorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0589.n-ary-tree-preorder-traversal/dfs/stack/main.go)| dfs.stack
| 0590 | [N-ary Tree Postorder Traversal  ](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal) | 76.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0590.n-ary-tree-postorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0590.n-ary-tree-postorder-traversal/dfs/stack/main.go)| dfs.stack
| 0606 | [Construct String from Binary Tree](https://leetcode-cn.com/problems/construct-string-from-binary-tree) | 56.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0606.construct-string-from-binary-tree/dfs/main.go)| dfs
| 0617 | [Merge Two Binary Trees          ](https://leetcode-cn.com/problems/merge-two-binary-trees) | 78.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0617.merge-two-binary-trees/dfs/recursive/main.go)| dfs.recursive
| 0653 | [Two Sum IV - Input is a BST     ](https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst) | 59.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0653.two-sum-iv-input-is-a-bst/dfs/main.go)| dfs
| 0654 | [Maximum Binary Tree             ](https://leetcode-cn.com/problems/maximum-binary-tree) | 81.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0654.maximum-binary-tree/dfs/main.go)| dfs
| 0669 | [Trim a Binary Search Tree       ](https://leetcode-cn.com/problems/trim-a-binary-search-tree) | 66.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0669.trim-a-binary-search-tree/dfs/main.go)| dfs
| 0671 | [Second Minimum Node In a Binary Tree](https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree) | 48.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0671.second-minimum-node-in-a-binary-tree/dfs/main.go)| dfs
| 0687 | [Longest Univalue Path           ](https://leetcode-cn.com/problems/longest-univalue-path) | 44.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0687.longest-univalue-path/dfs/main.go)| dfs
| 0700 | [Search in a Binary Search Tree  ](https://leetcode-cn.com/problems/search-in-a-binary-search-tree) | 75.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0700.search-in-a-binary-search-tree/dfs/main.go)| dfs
| 0701 | [Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree) | 72.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0701.insert-into-a-binary-search-tree/dfs/main.go)| dfs
| 0783 | [Minimum Distance Between BST Nodes](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes) | 59.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0783.minimum-distance-between-bst-nodes/dfs/main.go)| dfs
| 0814 | [Binary Tree Pruning             ](https://leetcode-cn.com/problems/binary-tree-pruning) | 70.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0814.binary-tree-pruning/dfs/main.go)| dfs
| 0865 | [Smallest Subtree with all the Deepest Nodes](https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes) | 65.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0865.smallest-subtree-with-all-the-deepest-nodes/main.go)| 
| 0872 | [Leaf-Similar Trees              ](https://leetcode-cn.com/problems/leaf-similar-trees) | 65.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0872.leaf-similar-trees/dfs/main.go)| dfs
| 0889 | [Construct Binary Tree from Preorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal) | 67.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0889.construct-binary-tree-from-preorder-and-postorder-traversal/main.go)| 
| 0897 | [Increasing Order Search Tree    ](https://leetcode-cn.com/problems/increasing-order-search-tree) | 74.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0897.increasing-order-search-tree/inorder/main.go)| inorder
| 0938 | [Range Sum of BST                ](https://leetcode-cn.com/problems/range-sum-of-bst) | 81.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0938.range-sum-of-bst/main.go)| 
| 0951 | [Flip Equivalent Binary Trees    ](https://leetcode-cn.com/problems/flip-equivalent-binary-trees) | 66.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0951.flip-equivalent-binary-trees/dfs/main.go)| dfs
| 0965 | [Univalued Binary Tree           ](https://leetcode-cn.com/problems/univalued-binary-tree) | 68.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0965.univalued-binary-tree/dfs/main.go)| dfs
| 0979 | [Distribute Coins in Binary Tree ](https://leetcode-cn.com/problems/distribute-coins-in-binary-tree) | 71.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0979.distribute-coins-in-binary-tree/dfs/main.go)| dfs
| 1008 | [Construct Binary Search Tree from Preorder Traversal](https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal) | 72.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1008.construct-binary-search-tree-from-preorder-traversal/dfs/main.go)| dfs
| 1022 | [Sum of Root To Leaf Binary Numbers](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers) | 70.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/1022.sum-of-root-to-leaf-binary-numbers/dfs/main.go)| dfs
| 1026 | [Maximum Difference Between Node and Ancestor](https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor) | 66.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1026.maximum-difference-between-node-and-ancestor/dfs/main.go)| dfs
| 1038 | [Binary Search Tree to Greater Sum Tree](https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree) | 79.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1038.binary-search-tree-to-greater-sum-tree/inorder/main.go)| inorder
| 1104 | [Path In Zigzag Labelled Binary Tree](https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree) | 76.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1104.path-in-zigzag-labelled-binary-tree/main.go)| 
| 1123 | [Lowest Common Ancestor of Deepest Leaves](https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves) | 70.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1123.lowest-common-ancestor-of-deepest-leaves/dfs/main.go)| dfs
| 1161 | [Maximum Level Sum of a Binary Tree](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree) | 63.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1161.maximum-level-sum-of-a-binary-tree/bfs/main.go)| bfs
| 1302 | [Deepest Leaves Sum              ](https://leetcode-cn.com/problems/deepest-leaves-sum) | 81.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1302.deepest-leaves-sum/dfs/main.go)| dfs
| 1315 | [Sum of Nodes with Even-Valued Grandparent](https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent) | 81.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1315.sum-of-nodes-with-even-valued-grandparent/dfs/main.go)| dfs
| 1325 | [Delete Leaves With a Given Value](https://leetcode-cn.com/problems/delete-leaves-with-a-given-value) | 72.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1325.delete-leaves-with-a-given-value/postorder/main.go)| postorder
| 1448 | [Count Good Nodes in Binary Tree ](https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree) | 71.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1448.count-good-nodes-in-binary-tree/dfs/main.go)| dfs
| 1484 | [Clone Binary Tree With Random Pointer](https://leetcode-cn.com/problems/clone-binary-tree-with-random-pointer) | 80.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1484.clone-binary-tree-with-random-pointer/dfs/main.go)| dfs
| 1490 | [Clone N-ary Tree                ](https://leetcode-cn.com/problems/clone-n-ary-tree) | 85.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1490.clone-n-ary-tree/dfs/main.go)| dfs
| 1560 | [Most Visited Sector in  a Circular Track](https://leetcode-cn.com/problems/most-visited-sector-in-a-circular-track) | 57.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/1560.most-visited-sector-in-a-circular-track/hash/main.go)| hash
| 1600 | [Throne Inheritance              ](https://leetcode-cn.com/problems/throne-inheritance) | 68.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1600.throne-inheritance%20皇位继承顺序/main.go)| 
| 1602 | [Find Nearest Right Node in Binary Tree](https://leetcode-cn.com/problems/find-nearest-right-node-in-binary-tree) | 77.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1602.find-nearest-right-node-in-binary-tree/bfs/main.go)| bfs
| 1669 | [Merge In Between Linked Lists   ](https://leetcode-cn.com/problems/merge-in-between-linked-lists) | 76.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1669.merge-in-between-linked-lists/main.go)| 
| 1740 | [Find Distance in a Binary Tree  ](https://leetcode-cn.com/problems/find-distance-in-a-binary-tree) | 70.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1740.find-distance-in-a-binary-tree/dfs/main.go)| dfs

## Divide&conquer
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0014 | [Longest Common Prefix           ](https://leetcode-cn.com/problems/longest-common-prefix) | 40.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/divide&conquer/0014.longest-common-prefix%20最长公共前缀/main.go)| 

## Back
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0037 | [Sudoku Solver                   ](https://leetcode-cn.com/problems/sudoku-solver) | 67.1% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/back/0037.sudoku-solver%20解数独/main.go)| 
| 0000 | [Offer                           ](https://leetcode-cn.com/problems/0038) | NaN% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/back/Offer.0038.zi-fu-chuan-de-pai-lie-lcof%20字符串的排列/main.go)| 

## Sort
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0179 | [Largest Number                  ](https://leetcode-cn.com/problems/largest-number) | 41.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sort/0179.largest-number%20最大数/main.go)| 

## Search
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0011 | [Container With Most Water       ](https://leetcode-cn.com/problems/container-with-most-water) | 63.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/0011.container-with-most-water%20盛最多水的容器/main.go)| 
| 0015 | [3Sum                            ](https://leetcode-cn.com/problems/3sum) | 33.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/0015.3sum%20三数之和/main.go)| 
| 0016 | [3Sum Closest                    ](https://leetcode-cn.com/problems/3sum-closest) | 46.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/0016.3sum-closest%20最接近的三数之和/main.go)| 
| 0018 | [4Sum                            ](https://leetcode-cn.com/problems/4sum) | 40.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/0018.4sum%20四数之和/main.go)| 
| 0026 | [Remove Duplicates from Sorted Array](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array) | 54.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/search/0026.remove-duplicates-from-sorted-array%20删除有序数组中的重复项/main.go)| 
| 1838 | [Frequency of the Most Frequent Element](https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element) | 43.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/1838.frequency-of-the-most-frequent-element%20最高频元素的频数/main.go)| 
| 0000 | [Offer 53 - I                    ](https://leetcode-cn.com/problems/ 在排序数组中查找数字 I) | NaN% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/search/Offer%2053%20-%20I.%20在排序数组中查找数字%20I/main.go)| 

## LinkedList
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0021 | [Merge Two Sorted Lists          ](https://leetcode-cn.com/problems/merge-two-sorted-lists) | 66.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0021.merge-two-sorted-lists%20合并两个有序链表/main.go)| 
| 0023 | [Merge k Sorted Lists            ](https://leetcode-cn.com/problems/merge-k-sorted-lists) | 55.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0023.merge-k-sorted-lists%20合并K个有序链表/main.go)| 
| 0024 | [Swap Nodes in Pairs             ](https://leetcode-cn.com/problems/swap-nodes-in-pairs) | 70.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0024.swap-nodes-in-pairs%20两两交换节点/main.go)| 
| 0025 | [Reverse Nodes in k-Group        ](https://leetcode-cn.com/problems/reverse-nodes-in-k-group) | 65.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0025.reverse-nodes-in-k-group%20K%20个一组翻转链表/main.go)| 
| 0061 | [Rotate List                     ](https://leetcode-cn.com/problems/rotate-list) | 41.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0061.rotate-list%20拼接链表/main.go)| 
| 0082 | [Remove Duplicates from Sorted List II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii) | 52.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0082.remove-duplicates-from-sorted-list-ii%20删除重复全部元素/main.go)| 
| 0083 | [Remove Duplicates from Sorted List](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list) | 53.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0083.remove-duplicates-from-sorted-list%20删除重复多余元素/main.go)| 
| 0086 | [Partition List                  ](https://leetcode-cn.com/problems/partition-list) | 63.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0086.partition-list%20小于X的靠左/main.go)| 
| 0092 | [Reverse Linked List II          ](https://leetcode-cn.com/problems/reverse-linked-list-ii) | 54.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0092.reverse-linked-list-ii%20反转链表/main.go)| 
| 0109 | [Convert Sorted List to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree) | 76.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0109.convert-sorted-list-to-binary-search-tree%20有序链表转二叉搜索树/array/main.go)| array
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0109.convert-sorted-list-to-binary-search-tree%20有序链表转二叉搜索树/inorder/main.go)| inorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0109.convert-sorted-list-to-binary-search-tree%20有序链表转二叉搜索树/recursive/main.go)| recursive
| 0138 | [Copy List with Random Pointer   ](https://leetcode-cn.com/problems/copy-list-with-random-pointer) | 65.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0138.copy-list-with-random-pointer%20深拷贝随机指针链表/main.go)| 
| 0141 | [Linked List Cycle               ](https://leetcode-cn.com/problems/linked-list-cycle) | 51.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0141.linked-list-cycle%20是否有环/main.go)| 
| 0142 | [Linked List Cycle II            ](https://leetcode-cn.com/problems/linked-list-cycle-ii) | 55.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0142.linked-list-cycle-ii%20环的起点/main.go)| 
| 0143 | [Reorder List                    ](https://leetcode-cn.com/problems/reorder-list) | 61.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0143.reorder-list%20重排链表/main.go)| 
| 0147 | [Insertion Sort List             ](https://leetcode-cn.com/problems/insertion-sort-list) | 68.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0147.insertion-sort-list%20插入排序/main.go)| 
| 0160 | [Intersection of Two Linked Lists](https://leetcode-cn.com/problems/intersection-of-two-linked-lists) | 61.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0160.intersection-of-two-linked-lists%20相交起始点/main.go)| 
| 0203 | [Remove Linked List Elements     ](https://leetcode-cn.com/problems/remove-linked-list-elements) | 51.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0203.remove-linked-list-elements%20删除链表元素/main.go)| 
| 0206 | [Reverse Linked List             ](https://leetcode-cn.com/problems/reverse-linked-list) | 72.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0206.reverse-linked-list%20反转链表/main.go)| 
| 0234 | [Palindrome Linked List          ](https://leetcode-cn.com/problems/palindrome-linked-list) | 49.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0234.palindrome-linked-list%20回文链表/main.go)| 
| 0237 | [Delete Node in a Linked List    ](https://leetcode-cn.com/problems/delete-node-in-a-linked-list) | 84.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0237.delete-node-in-a-linked-list%20删除节点/main.go)| 
| 0328 | [Odd Even Linked List            ](https://leetcode-cn.com/problems/odd-even-linked-list) | 65.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0328.odd-even-linked-list%20奇偶链表/main.go)| 
| 0369 | [Plus One Linked List            ](https://leetcode-cn.com/problems/plus-one-linked-list) | 61.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0369.plus-one-linked-list%20加一/main.go)| 
| 0379 | [Design Phone Directory          ](https://leetcode-cn.com/problems/design-phone-directory) | 64.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0379.design-phone-directory/main.go)| 
| 0430 | [Flatten a Multilevel Doubly Linked List](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list) | 54.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0430.flatten-a-multilevel-doubly-linked-list%20扁平化多级双向链表/main.go)| 
| 0445 | [Add Two Numbers II              ](https://leetcode-cn.com/problems/add-two-numbers-ii) | 58.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0445.add-two-numbers-ii%20两数相加/main.go)| 
| 0707 | [Design Linked List              ](https://leetcode-cn.com/problems/design-linked-list) | 31.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0707.design-linked-list%20设计链表/main.go)| 
| 0708 | [Insert into a Sorted Circular Linked List](https://leetcode-cn.com/problems/insert-into-a-sorted-circular-linked-list) | 33.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0708.insert-into-a-sorted-circular-linked-list%20循环有序列表的插入/main.go)| 
| 0725 | [Split Linked List in Parts      ](https://leetcode-cn.com/problems/split-linked-list-in-parts) | 56.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0725.split-linked-list-in-parts%20分割链表/main.go)| 
| 0817 | [Linked List Components          ](https://leetcode-cn.com/problems/linked-list-components) | 59.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0817.linked-list-components%20链表组件/main.go)| 
| 0876 | [Middle of the Linked List       ](https://leetcode-cn.com/problems/middle-of-the-linked-list) | 70.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0876.middle-of-the-linked-list%20中间节点/main.go)| 
| 1019 | [Next Greater Node In Linked List](https://leetcode-cn.com/problems/next-greater-node-in-linked-list) | 59.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1019.next-greater-node-in-linked-list%20下一个更大的值/main.go)| 
| 1171 | [Remove Zero Sum Consecutive Nodes from Linked List](https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list) | 47.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1171.remove-zero-sum-consecutive-nodes-from-linked-list%20从链表中删去总和值为零的连续节点/main.go)| 
| 1290 | [Convert Binary Number in a Linked List to Integer](https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer) | 80.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1290.convert-binary-number-in-a-linked-list-to-integer%20二进制链表的值/main.go)| 
| 1367 | [Linked List in Binary Tree      ](https://leetcode-cn.com/problems/linked-list-in-binary-tree) | 41.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1367.linked-list-in-binary-tree%20二叉树的链表/main.go)| 
| 1474 | [Delete N Nodes After M Nodes of a Linked List](https://leetcode-cn.com/problems/delete-n-nodes-after-m-nodes-of-a-linked-list) | 72.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1474.delete-n-nodes-after-m-nodes-of-a-linked-list%20%20删除m节点后移动n节点/main.go)| 
| 1634 | [Add Two Polynomials Represented as Linked Lists](https://leetcode-cn.com/problems/add-two-polynomials-represented-as-linked-lists) | 52.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1634.add-two-polynomials-represented-as-linked-lists%20多项式合并/main.go)| 
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1634.add-two-polynomials-represented-as-linked-lists%20多项式合并/main.js)| 
| 1669 | [Merge In Between Linked Lists   ](https://leetcode-cn.com/problems/merge-in-between-linked-lists) | 76.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1669.merge-in-between-linked-lists%20删除区间后合并/main.go)| 
| 1670 | [Design Front Middle Back Queue  ](https://leetcode-cn.com/problems/design-front-middle-back-queue) | 52.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1670.design-front-middle-back-queue%20队列/main.go)| 
| 1721 | [Swapping Nodes in a Linked List ](https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list) | 64.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1721.swapping-nodes-in-a-linked-list%20交换镜像节点/main.go)| 

## BinarySearch
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0004 | [Median of Two Sorted Arrays     ](https://leetcode-cn.com/problems/median-of-two-sorted-arrays) | 40.7% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0004.median-of-two-sorted-arrays%20寻找两个正序数组的中位数/main.go)| 
| 0014 | [Longest Common Prefix           ](https://leetcode-cn.com/problems/longest-common-prefix) | 40.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0014.longest-common-prefix%20最长公共前缀/main.go)| 
| 0033 | [Search in Rotated Sorted Array  ](https://leetcode-cn.com/problems/search-in-rotated-sorted-array) | 42.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0033.search-in-rotated-sorted-array%20搜索旋转排序数组/main.go)| 
| 0034 | [Find First and Last Position of Element in Sorted Array](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array) | 42.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0034.find-first-and-last-position-of-element-in-sorted-array%20在排序数组中查找元素的第一个和最后一个位置/main.go)| 
| 0035 | [Search Insert Position          ](https://leetcode-cn.com/problems/search-insert-position) | 46.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0035.search-insert-position%20搜索插入位置/main.go)| 
| 0278 | [First Bad Version               ](https://leetcode-cn.com/problems/first-bad-version) | 45.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0278.first-bad-version%20第一个错误版本/main.go)| 
| 0374 | [Guess Number Higher or Lower    ](https://leetcode-cn.com/problems/guess-number-higher-or-lower) | 51.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0374.guess-number-higher-or-lower%20猜数字大小/main.go)| 
| 0701 | [Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree) | 72.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0701.insert-into-a-binary-search-tree%20二叉搜索树中的插入操作/main.go)| 
| 0852 | [Peak Index in a Mountain Array  ](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array) | 71.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0852.peak-index-in-a-mountain-array%20山脉数组的峰顶索引/main.go)| 

## Catalan
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0022 | [Generate Parentheses            ](https://leetcode-cn.com/problems/generate-parentheses) | 77.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/catalan/0022.generate-parentheses/main.go)| 

## StackMonotonic
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0042 | [Trapping Rain Water             ](https://leetcode-cn.com/problems/trapping-rain-water) | 57.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/stack_monotonic/0042.trapping-rain-water%20接雨水/main.go)| 

## SortCounter
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0274 | [H-Index                         ](https://leetcode-cn.com/problems/h-index) | 44.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sort_counter/0274.h-index%20H%20指数/main.go)| 
| 0451 | [Sort Characters By Frequency    ](https://leetcode-cn.com/problems/sort-characters-by-frequency) | 71.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sort_counter/0451.sort-characters-by-frequency%20根据字符出现频率排序/main.go)| 
| 1365 | [How Many Numbers Are Smaller Than the Current Number](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number) | 82.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sort_counter/1365.how-many-numbers-are-smaller-than-the-current-number%20有多少小于当前数字的数字/main.go)| 
| 1833 | [Maximum Ice Cream Bars          ](https://leetcode-cn.com/problems/maximum-ice-cream-bars) | 68.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sort_counter/1833.maximum-ice-cream-bars%20雪糕的最大数量/main.go)| 

## Hash
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0001 | [Two Sum                         ](https://leetcode-cn.com/problems/two-sum) | 51.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0001.two-sum%20两数之和/main.go)| 
| 0003 | [Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) | 37.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/0003.longest-substring-without-repeating-characters%20无重复的最长子串/main.go)| 
| 0041 | [First Missing Positive          ](https://leetcode-cn.com/problems/first-missing-positive) | 41.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/hash/0041.first-missing-positive%20缺失的第一个正数/main.go)| 
| 0170 | [Two Sum III - Data structure design](https://leetcode-cn.com/problems/two-sum-iii-data-structure-design) | 41.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0170.two-sum-iii-data-structure-design%20两数之和%20III%20-%20数据结构设计/main.go)| 
| 0205 | [Isomorphic Strings              ](https://leetcode-cn.com/problems/isomorphic-strings) | 50.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0205.isomorphic-strings%20同构字符串/main.go)| 
| 0217 | [Contains Duplicate              ](https://leetcode-cn.com/problems/contains-duplicate) | 56.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0217.contains-duplicate%20出现重复/main.go)| 
| 0219 | [Contains Duplicate II           ](https://leetcode-cn.com/problems/contains-duplicate-ii) | 42.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0219.contains-duplicate-ii%20两个重复元素距离小于等于K/main.go)| 
| 0242 | [Valid Anagram                   ](https://leetcode-cn.com/problems/valid-anagram) | 64.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0242.valid-anagram%20有效的字母异位词/main.go)| 
| 0246 | [Strobogrammatic Number          ](https://leetcode-cn.com/problems/strobogrammatic-number) | 47.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0246.strobogrammatic-number%20中心对称数/main.go)| 
| 0266 | [Palindrome Permutation          ](https://leetcode-cn.com/problems/palindrome-permutation) | 67.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0266.palindrome-permutation%20回文排列/main.go)| 
| 0290 | [Word Pattern                    ](https://leetcode-cn.com/problems/word-pattern) | 45.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0290.word-pattern%20单词规律/main.go)| 
| 0349 | [Intersection of Two Arrays      ](https://leetcode-cn.com/problems/intersection-of-two-arrays) | 73.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0349.intersection-of-two-arrays%20两个数组的交集/main.go)| 
| 0350 | [Intersection of Two Arrays II   ](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii) | 55.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0350.intersection-of-two-arrays-ii%20两个数组的交集%20II/main.go)| 
| 0359 | [Logger Rate Limiter             ](https://leetcode-cn.com/problems/logger-rate-limiter) | 72.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0359.logger-rate-limiter%20日志速率限制器/main.go)| 
| 0383 | [Ransom Note                     ](https://leetcode-cn.com/problems/ransom-note) | 59.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0383.ransom-note%20赎金信/main.go)| 
| 0387 | [First Unique Character in a String](https://leetcode-cn.com/problems/first-unique-character-in-a-string) | 53.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0387.first-unique-character-in-a-string%20字符串中的第一个唯一字符/main.go)| 
| 0389 | [Find the Difference             ](https://leetcode-cn.com/problems/find-the-difference) | 69.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0389.find-the-difference%20找不同/main.go)| 
| 0409 | [Longest Palindrome              ](https://leetcode-cn.com/problems/longest-palindrome) | 55.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0409.longest-palindrome%20最长回文串/main.go)| 
| 0448 | [Find All Numbers Disappeared in an Array](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array) | 64.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0448.find-all-numbers-disappeared-in-an-array%20找到所有数组中消失的数字/main.go)| 
| 0500 | [Keyboard Row                    ](https://leetcode-cn.com/problems/keyboard-row) | 70.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0500.keyboard-row%20只返回键盘同一行的字母/main.go)| 
| 0575 | [Distribute Candies              ](https://leetcode-cn.com/problems/distribute-candies) | 68.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0575.distribute-candies%20分糖果/main.go)| 
| 0594 | [Longest Harmonious Subsequence  ](https://leetcode-cn.com/problems/longest-harmonious-subsequence) | 51.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0594.longest-harmonious-subsequence%20最长和谐(相差1)子序列/main.go)| 
| 0599 | [Minimum Index Sum of Two Lists  ](https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists) | 52.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0599.minimum-index-sum-of-two-lists%20相同值的最小索引总和/main.go)| 
| 0604 | [Design Compressed String Iterator](https://leetcode-cn.com/problems/design-compressed-string-iterator) | 37.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0604.design-compressed-string-iterator%20迭代压缩字符串/main.go)| 
| 0657 | [Robot Return to Origin          ](https://leetcode-cn.com/problems/robot-return-to-origin) | 78.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0657.robot-return-to-origin%20机器人能否返回原点/main.go)| 
| 0690 | [Employee Importance             ](https://leetcode-cn.com/problems/employee-importance) | 64.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0690.employee-importance%20员工的重要性/main.go)| 
| 0697 | [Degree of an Array              ](https://leetcode-cn.com/problems/degree-of-an-array) | 60.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0697.degree-of-an-array%20数组的度/main.go)| 
| 0705 | [Design HashSet                  ](https://leetcode-cn.com/problems/design-hashset) | 64.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0705.design-hashset%20设计哈希集合/main.go)| 
| 0706 | [Design HashMap                  ](https://leetcode-cn.com/problems/design-hashmap) | 64.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0706.design-hashmap%20设计哈希映射/main.go)| 
| 0720 | [Longest Word in Dictionary      ](https://leetcode-cn.com/problems/longest-word-in-dictionary) | 48.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0720.longest-word-in-dictionary%20词典中最长的单词/main.go)| 
| 0734 | [Sentence Similarity             ](https://leetcode-cn.com/problems/sentence-similarity) | 47.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0734.sentence-similarity%20句子相似性/main.go)| 
| 0748 | [Shortest Completing Word        ](https://leetcode-cn.com/problems/shortest-completing-word) | 59.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0748.shortest-completing-word%20最短补全词/main.go)| 
| 0760 | [Find Anagram Mappings           ](https://leetcode-cn.com/problems/find-anagram-mappings) | 83.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0760.find-anagram-mappings%20找A在B的位置/main.go)| 
| 0771 | [Jewels and Stones               ](https://leetcode-cn.com/problems/jewels-and-stones) | 85.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0771.jewels-and-stones%20找J在S中的数量/main.go)| 
| 0811 | [Subdomain Visit Count           ](https://leetcode-cn.com/problems/subdomain-visit-count) | 70.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/0811.subdomain-visit-count%20子域名访问次数/main.go)| 
| 0819 | [Most Common Word                ](https://leetcode-cn.com/problems/most-common-word) | 42.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0819.most-common-word%20最常见的单词/main.go)| 
| 0859 | [Buddy Strings                   ](https://leetcode-cn.com/problems/buddy-strings) | 30.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0859.buddy-strings%20亲密字符串/main.go)| 
| 0884 | [Uncommon Words from Two Sentences](https://leetcode-cn.com/problems/uncommon-words-from-two-sentences) | 66.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0884.uncommon-words-from-two-sentences%20两句话中的不常见单词/main.go)| 
| 0888 | [Fair Candy Swap                 ](https://leetcode-cn.com/problems/fair-candy-swap) | 64.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0888.fair-candy-swap%20公平的糖果棒交换/main.go)| 
| 0914 | [X of a Kind in a Deck of Cards  ](https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards) | 38.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0914.x-of-a-kind-in-a-deck-of-cards%20卡牌分组/main.go)| 
| 0929 | [Unique Email Addresses          ](https://leetcode-cn.com/problems/unique-email-addresses) | 64.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0929.unique-email-addresses%20独特的电子邮件地址/main.go)| 
| 0930 | [Binary Subarrays With Sum       ](https://leetcode-cn.com/problems/binary-subarrays-with-sum) | 54.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/0930.binary-subarrays-with-sum%20和相同的二元子数组/main.go)| 
| 0953 | [Verifying an Alien Dictionary   ](https://leetcode-cn.com/problems/verifying-an-alien-dictionary) | 55.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0953.verifying-an-alien-dictionary%20验证外星语词典/main.go)| 
| 0961 | [N-Repeated Element in Size 2N Array](https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array) | 67.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0961.n-repeated-element-in-size-2n-array%20重复%20N%20次的元素/main.go)| 
| 0997 | [Find the Town Judge             ](https://leetcode-cn.com/problems/find-the-town-judge) | 51.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0997.find-the-town-judge%20找到小镇的法官/main.go)| 
| 1002 | [Find Common Characters          ](https://leetcode-cn.com/problems/find-common-characters) | 73.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1002.find-common-characters%20查找常用字符/main.go)| 
| 1086 | [High Five                       ](https://leetcode-cn.com/problems/high-five) | 68.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1086.high-five%20前五科的均分/main.go)| 
| 1119 | [Remove Vowels from a String     ](https://leetcode-cn.com/problems/remove-vowels-from-a-string) | 86.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1119.remove-vowels-from-a-string%20删去字符串中的元音/main.go)| 
| 1122 | [Relative Sort Array             ](https://leetcode-cn.com/problems/relative-sort-array) | 70.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1122.relative-sort-array%20数组的相对排序/main.go)| 
| 1128 | [Number of Equivalent Domino Pairs](https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs) | 54.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1128.number-of-equivalent-domino-pairs%20等价多米诺骨牌对的数量/main.go)| 
| 1133 | [Largest Unique Number           ](https://leetcode-cn.com/problems/largest-unique-number) | 64.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1133.largest-unique-number%20最大唯一数/main.go)| 
| 1160 | [Find Words That Can Be Formed by Characters](https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters) | 68.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1160.find-words-that-can-be-formed-by-characters%20拼写单词/main.go)| 
| 1165 | [Single-Row Keyboard             ](https://leetcode-cn.com/problems/single-row-keyboard) | 83.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1165.single-row-keyboard%20机械手最少移动次数/main.go)| 
| 1331 | [Rank Transform of an Array      ](https://leetcode-cn.com/problems/rank-transform-of-an-array) | 53.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1331.rank-transform-of-an-array%20数组序号转换/main.go)| 
| 1436 | [Destination City                ](https://leetcode-cn.com/problems/destination-city) | 78.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1436.destination-city%20旅行终点站/main.go)| 
| 1512 | [Number of Good Pairs            ](https://leetcode-cn.com/problems/number-of-good-pairs) | 84.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1512.number-of-good-pairs/main.go)| 
| 1570 | [Dot Product of Two Sparse Vectors](https://leetcode-cn.com/problems/dot-product-of-two-sparse-vectors) | 90.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/1570.dot-product-of-two-sparse-vectors/main.go)| 
| 1603 | [Design Parking System           ](https://leetcode-cn.com/problems/design-parking-system) | 84.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1603.design-parking-system/main.go)| 
| 1711 | [Count Good Meals                ](https://leetcode-cn.com/problems/count-good-meals) | 35.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/1711.count-good-meals%20大餐计数/main.go)| 
| 1748 | [Sum of Unique Elements          ](https://leetcode-cn.com/problems/sum-of-unique-elements) | 75.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1748.sum-of-unique-elements%20唯一元素的和/main.go)| 
| 1805 | [Number of Different Integers in a String](https://leetcode-cn.com/problems/number-of-different-integers-in-a-string) | 49.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1805.number-of-different-integers-in-a-string%20字符串中不同整数的数目/main.go)| 
| 0000 | [LCP                             ](https://leetcode-cn.com/problems/07) | NaN% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/hash/LCP.07.%20传递信息/main.go)| 

## Greedy
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0045 | [Jump Game II                    ](https://leetcode-cn.com/problems/jump-game-ii) | 42.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/greedy/0045.jump-game-ii%20跳跃游戏%20II/main.go)| 
| 0055 | [Jump Game                       ](https://leetcode-cn.com/problems/jump-game) | 43.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/greedy/0055.jump-game%20跳跃游戏/main.go)| 

## Dp
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0005 | [Longest Palindromic Substring   ](https://leetcode-cn.com/problems/longest-palindromic-substring) | 35.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0005.longest-palindromic-substring%20最长回文/main.go)| 
| 0010 | [Regular Expression Matching     ](https://leetcode-cn.com/problems/regular-expression-matching) | 31.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0010.regular-expression-matching%20.*正则匹配/main.go)| 
| 0032 | [Longest Valid Parentheses       ](https://leetcode-cn.com/problems/longest-valid-parentheses) | 35.6% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0032.longest-valid-parentheses%20最长有效括号/main.go)| 
| 0042 | [Trapping Rain Water             ](https://leetcode-cn.com/problems/trapping-rain-water) | 57.5% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0042.trapping-rain-water%20接雨水/main.go)| 
| 0044 | [Wildcard Matching               ](https://leetcode-cn.com/problems/wildcard-matching) | 32.8% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0044.wildcard-matching%20*？通配符/main.go)| 
| 0053 | [Maximum Subarray                ](https://leetcode-cn.com/problems/maximum-subarray) | 55.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0053.maximum-subarray%20最大子序和/main.go)| 
| 0062 | [Unique Paths                    ](https://leetcode-cn.com/problems/unique-paths) | 66.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0062.unique-paths%20不同路径/main.go)| 
| 0063 | [Unique Paths II                 ](https://leetcode-cn.com/problems/unique-paths-ii) | 39.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0063.unique-paths-ii%20不同路径%20II/main.go)| 
| 0064 | [Minimum Path Sum                ](https://leetcode-cn.com/problems/minimum-path-sum) | 68.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0064.minimum-path-sum%20最小路径和/main.go)| 
| 0070 | [Climbing Stairs                 ](https://leetcode-cn.com/problems/climbing-stairs) | 52.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0070.climbing-stairs%20爬楼梯/main.go)| 
| 0072 | [Edit Distance                   ](https://leetcode-cn.com/problems/edit-distance) | 61.3% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0072.edit-distance%20编辑距离/main.go)| 
| 0087 | [Scramble String                 ](https://leetcode-cn.com/problems/scramble-string) | 48.9% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0087.scramble-string%20扰乱字符串/main.go)| 
| 0091 | [Decode Ways                     ](https://leetcode-cn.com/problems/decode-ways) | 30.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0091.decode-ways%20解码方法/main.go)| 
| 0097 | [Interleaving String             ](https://leetcode-cn.com/problems/interleaving-string) | 45.8% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0097.interleaving-string%20交错字符串/main.go)| 
| 0115 | [Distinct Subsequences           ](https://leetcode-cn.com/problems/distinct-subsequences) | 52.4% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0115.distinct-subsequences%20不同的子序列/main.go)| 
| 0118 | [Pascal's Triangle               ](https://leetcode-cn.com/problems/pascals-triangle) | 72.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0118.pascals-triangle%20杨辉三角/main.go)| 
| 0119 | [Pascal's Triangle II            ](https://leetcode-cn.com/problems/pascals-triangle-ii) | 66.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0119.pascals-triangle-ii%20杨辉三角%20II/main.go)| 
| 0120 | [Triangle                        ](https://leetcode-cn.com/problems/triangle) | 67.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0120.triangle%20三角形最小路径/main.go)| 
| 0121 | [Best Time to Buy and Sell Stock ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock) | 57.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0121.best-time-to-buy-and-sell-stock%20买卖股票的最佳时机/main.go)| 
| 0123 | [Best Time to Buy and Sell Stock III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii) | 53.4% | lHard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0123.best-time-to-buy-and-sell-stock-iii%20买卖股票的最佳时机%20III/main.go)| 
| 0139 | [Word Break                      ](https://leetcode-cn.com/problems/word-break) | 51.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0139.word-break%20单词拆分/main.go)| 
| 0279 | [Perfect Squares                 ](https://leetcode-cn.com/problems/perfect-squares) | 63.1% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0279.perfect-squares%20完全平方数/main.go)| 
| 0303 | [Range Sum Query - Immutable     ](https://leetcode-cn.com/problems/range-sum-query-immutable) | 72.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0303.range-sum-query-immutable%20区域和检索%20-%20数组不可变/main.go)| 
| 0322 | [Coin Change                     ](https://leetcode-cn.com/problems/coin-change) | 44.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0322.coin-change%20零钱兑换/main.go)| 
| 0474 | [Ones and Zeroes                 ](https://leetcode-cn.com/problems/ones-and-zeroes) | 61.0% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0474.ones-and-zeroes%20一和零/main.go)| 
| 0509 | [Fibonacci Number                ](https://leetcode-cn.com/problems/fibonacci-number) | 67.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0509.fibonacci-number%20斐波那契数/main.go)| 
| 0518 | [Coin Change 2                   ](https://leetcode-cn.com/problems/coin-change-2) | 65.2% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0518.coin-change-2%20零钱兑换%20II/main.go)| 
| 0523 | [Continuous Subarray Sum         ](https://leetcode-cn.com/problems/continuous-subarray-sum) | 27.4% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0523.continuous-subarray-sum%20连续子数组和%20K的倍数/main.go)| 
| 0525 | [Contiguous Array                ](https://leetcode-cn.com/problems/contiguous-array) | 53.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0525.contiguous-array%20连续数组/main.go)| 
| 0877 | [Stone Game                      ](https://leetcode-cn.com/problems/stone-game) | 75.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0877.stone-game%20石头游戏/main.go)| 
| 1480 | [Running Sum of 1d Array         ](https://leetcode-cn.com/problems/running-sum-of-1d-array) | 86.9% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/1480.running-sum-of-1d-array%20一维数组的动态和/main.go)| 
| 1744 | [Can You Eat Your Favorite Candy on Your Favorite Day?](https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day) | 36.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/1744.can-you-eat-your-favorite-candy-on-your-favorite-day%20你能在你最喜欢的那天吃到你最喜欢的糖果吗？/main.go)| 
| 1769 | [Minimum Number of Operations to Move All Balls to Each Box](https://leetcode-cn.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box) | 85.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/1769.minimum-number-of-operations-to-move-all-balls-to-each-box%20移动所有球到每个盒子所需的最小操作数/main.go)| 

## Classic
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0005 | [Longest Palindromic Substring   ](https://leetcode-cn.com/problems/longest-palindromic-substring) | 35.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/classic/0005.longest-palindromic-substring/manacher/main.go)| manacher
| 0028 | [Implement strStr()              ](https://leetcode-cn.com/problems/implement-strstr) | 40.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/classic/0028.implement-strstr/kmp/main.go)| kmp
| 0169 | [Majority Element                ](https://leetcode-cn.com/problems/majority-element) | 66.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/classic/0169.majority-element%20众数/Boyer-Moore%20投票/main.go)| Boyer-Moore 投票

## Math
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0007 | [Reverse Integer                 ](https://leetcode-cn.com/problems/reverse-integer) | 35.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0007.reverse-integer/main.go)| 
| 0008 | [String to Integer (atoi)        ](https://leetcode-cn.com/problems/string-to-integer-atoi) | 21.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0008.string-to-integer-atoi/main.go)| 
| 0009 | [Palindrome Number               ](https://leetcode-cn.com/problems/palindrome-number) | 58.5% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0009.palindrome-number/main.go)| 
| 0012 | [Integer to Roman                ](https://leetcode-cn.com/problems/integer-to-roman) | 66.5% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0012.integer-to-roman/main.go)| 
| 0013 | [Roman to Integer                ](https://leetcode-cn.com/problems/roman-to-integer) | 63.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0013.roman-to-integer/main.go)| 
| 0029 | [Divide Two Integers             ](https://leetcode-cn.com/problems/divide-two-integers) | 20.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0029.divide-two-integers/main.go)| 
| 0031 | [Next Permutation                ](https://leetcode-cn.com/problems/next-permutation) | 37.3% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0031.next-permutation/main.go)| 
| 0043 | [Multiply Strings                ](https://leetcode-cn.com/problems/multiply-strings) | 44.9% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0043.multiply-strings%20字符串相乘/main.go)| 
| 0066 | [Plus One                        ](https://leetcode-cn.com/problems/plus-one) | 45.8% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0066.plus-one/main.go)| 
| 0191 | [Number of 1 Bits                ](https://leetcode-cn.com/problems/number-of-1-bits) | 74.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0191.number-of-1-bits/main.go)| 
| 0268 | [Missing Number                  ](https://leetcode-cn.com/problems/missing-number) | 62.3% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0268.missing-number%20丢失的数字/main.go)| 
| 0292 | [Nim Game                        ](https://leetcode-cn.com/problems/nim-game) | 69.7% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0292.nim-game/main.go)| 
| 0319 | [Bulb Switcher                   ](https://leetcode-cn.com/problems/bulb-switcher) | 51.6% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0319.bulb-switcher/main.go)| 
| 0415 | [Add Strings                     ](https://leetcode-cn.com/problems/add-strings) | 53.6% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0415.add-strings%20字符串相加/main.go)| 
| 0877 | [Stone Game                      ](https://leetcode-cn.com/problems/stone-game) | 75.7% | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0877.stone-game%20石头游戏/main.go)| 
| 1025 | [Divisor Game                    ](https://leetcode-cn.com/problems/divisor-game) | 71.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1025.divisor-game%20除数博弈/main.go)| 
| 1266 | [Minimum Time Visiting All Points](https://leetcode-cn.com/problems/minimum-time-visiting-all-points) | 82.4% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1266.minimum-time-visiting-all-points/main.go)| 
| 1281 | [Subtract the Product and Sum of Digits of an Integer](https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer) | 83.2% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1281.subtract-the-product-and-sum-of-digits-of-an-integer/main.go)| 
| 1313 | [Decompress Run-Length Encoded List](https://leetcode-cn.com/problems/decompress-run-length-encoded-list) | 83.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1313.decompress-run-length-encoded-list/main.go)| 
| 1403 | [Minimum Subsequence in Non-Increasing Order](https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order) | 69.1% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1403.minimum-subsequence-in-non-increasing-order/main.go)| 
| 1791 | [Find Center of Star Graph       ](https://leetcode-cn.com/problems/find-center-of-star-graph) | 81.0% | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1791.find-center-of-star-graph/main.go)| 