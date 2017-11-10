# -*- coding:utf-8 -*-
u'''cdays-3-exercise-3.py
    @note: 使用全局变量和函式的递归调用
'''

global col
global row
global pos_diag
global nag_diag

def output():
    global count
    print row
    count += 1

def do_queen(i):
    for j in range(0, 8):
        # 在同一对角线上的所有点（设下标为(i,j)），
        # 要么(i+j)是常数，要么(i-j)是常数。
        # 若该列，正对角线，负对角线上都没有皇后，则放入i皇后
        #
        # nag_diag[i+j]，i表示行号，j表示列号，这样理解，从左上角开始，下移到行号的地方，然后
		# 再右移到列号的地方，那个点所在的对角线就记录在nag_diag[i+j]里
		# pos_diag[i-j+7],理解成pos_diag[i+7-j]，从左上角开始，下移至行号的地方，然后移到最右，
		# 接着左移列号个单位，那个点所在的对角线就记录在pos_diag[row+7-l]了
        if col[j] == 1 and pos_diag[i-j+7] == 1 and nag_diag[i+j] == 1:
            row[i] = j
            col[j] = 0
            pos_diag[i-j+7] = 0
            nag_diag[i+j] = 0
            if i < 7:
                do_queen(i+1)
            else:
                output()
            # 如果失败，则返回到前一行，返回后要立马恢复原来的状态。
            col[j] = 1
            pos_diag[i-j+7] = 1
            nag_diag[i+j] = 1

if __name__ == '__main__':
    col = []
    row = []
    pos_diag = []
    nag_diag = []
    count = 0
    for index in range(0, 8):
        col.append(1)
        row.append(0)
    for index in range(0, 15):
        pos_diag.append(1)
        nag_diag.append(1)
    do_queen(0)
print 'Totally have %d solutions!' % count
