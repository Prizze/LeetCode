# [1394. Find Lucky Integer in an Array](https://leetcode.com/problems/find-lucky-integer-in-an-array/description)
### Easy
Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

Return the largest lucky integer in the array. If there is no lucky integer return -1.
## Решение
Используем ```map``` для хранения числа-частоты. Выдаем наибольший ```key``` или ```-1```, если нет ```lucky integer```