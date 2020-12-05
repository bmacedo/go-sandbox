package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/**
https://adventofcode.com/2020/day/1

The Elves in accounting just need you to fix your expense report (your puzzle input);
apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply
those two numbers together.

Find the two entries that sum to 2020; what do you get if you multiply them together?
*/
func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	// look up pair of entries that sum up to 2020. Return their product
	result, err := multiplyTwoEntriesThatSumToK(input, 2020)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// look up triple of entries that sum to 2020. Return their product.
	result, err = multiplyThreeEntriesThatSumToK(input, 2020)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func multiplyTwoEntriesThatSumToK(input []int, k int) (int, error) {
	set := make(map[int]bool)

	for i := range input {
		set[input[i]] = true
	}

	for i := range set {
		if set[k-i] == true {
			return i * (k - i), nil
		}
	}
	return 0, fmt.Errorf("there are no pair of entries that sum to 2020")
}

func multiplyThreeEntriesThatSumToK(input []int, k int) (int, error) {
	set := make(map[int]bool)

	for i := range input {
		set[input[i]] = true
	}

	for i := range set {
		if subResult, err := multiplyTwoEntriesThatSumToK(input, k-i); err == nil {
			return i * subResult, nil
		}
	}
	return 0, fmt.Errorf("there are no triple of entries that sum to 2020")
}

func readInput() ([]int, error) {
	var input []int

	//txt := `
	//1721
	//979
	//366
	//299
	//675
	//1456
	//`

	txt := `
		1864
		1880
		1300
		1961
		1577
		1900
		1307
		1818
		1736
		1846
		1417
		1372
		1351
		1860
		1738
		1525
		1798
		1218
		1723
		1936
		1725
		1998
		1466
		1922
		1782
		1947
		1717
		1914
		1843
		1732
		1918
		814
		1771
		1712
		1804
		1213
		1859
		1820
		1793
		1870
		1993
		1787
		1824
		1849
		1646
		1489
		1348
		1978
		1628
		1781
		2002
		1297
		1829
		1596
		1819
		1313
		1413
		1726
		1449
		1810
		1295
		1679
		1358
		1949
		1644
		1825
		1891
		490
		1962
		1939
		1228
		1889
		1977
		1980
		1763
		1752
		1983
		1785
		1678
		2000
		1857
		1658
		1863
		1330
		1380
		1799
		1789
		1633
		1663
		296
		1985
		1117
		1239
		1854
		1960
		2004
		1940
		1876
		1739
		1858
		1283
		1423
		1982
		1836
		1451
		1840
		1347
		1652
		1695
		1210
		1861
		1199
		1346
		1786
		1814
		1958
		1853
		1974
		1917
		1308
		654
		1743
		1847
		1367
		1559
		1614
		1897
		2003
		1886
		1885
		1682
		1204
		1986
		1816
		1994
		1817
		1751
		1701
		1619
		1970
		816
		1852
		1832
		1631
		703
		1604
		1444
		1842
		1984
		1259
		1948
		1620
		1681
		1822
		1865
		1521
		1741
		1455
		1909
		1764
		261
		1464
		1905
		1325
		1766
		1749
		1292
		1874
		1267
		1269
		1969
		1991
		1219
		1345
		1976
		1369
		1942
		1388
		1776
		1629
		1987
		1684
		1813
		1203
		1965
		1729
		1930
		1609
		1801
		1402
		121
		1833
		1898
		1957
		1051
		1430
		1893
		1784
		1800
		1910
	`

	scanner := bufio.NewScanner(strings.NewReader(txt))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}

	if err := scanner.Err(); err != nil {
		return input, err
	}
	return input, nil
}
