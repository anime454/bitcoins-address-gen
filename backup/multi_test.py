def job(a):
    for i in range(1,4):
        time.sleep(0.5)
        print('งานที่ %d รอบที่ %d'%(a,i))

if(__name__=='__main__'):
    for j in range(1,6):
        p = mp.Process(target=job,args=(j,))
        p.start()
    print('สั่งงานไปแล้ว')
