# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/description/)

### <span style="color: #fac31d; background-color: #3c3c3c; padding : 7px 12px; border-radius: 15px;"> Medium</span> <br/>

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself. 

## Решение
Алгоритм основан на поразрядном сложении цифр и записи результата сразу в список. Каждое сложение формирует ```carry``` - перенос единицы в следующий разряд (как в сложении в столбик). Эта единица учитывается в сложении следюущего разряда.