package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type rulesAndPieces struct {
	rules  []string
	pieces []string
}

type piece struct {
	x int
	m int
	a int
	s int
}

type rulePart struct {
	member string
	op     string
	value  int
	target string
}

type rule struct {
	apply func(piece) string
	parts []rulePart
}

func splitRange(rts valueRange, value int) (valueRange, valueRange, valueRange) {
	if value < rts.min {
		return invalidValueRange, invalidValueRange, rts
	}
	if value > rts.max {
		return rts, invalidValueRange, invalidValueRange
	}
	return valueRange{rts.min, value - 1}, valueRange{value, value}, valueRange{value + 1, rts.max}
}

func (r rule) ranges(rg ruleRange) []ruleRange {
	rr := []ruleRange{}

	for _, part := range r.parts {
		if part.member == "" {
			newRR := rg
			newRR.step = part.target
			rr = append(rr, newRR)
			break
		}
		rangesToSplit := rg.ranges(part.member)
		var remainingRanges []valueRange
		for _, rts := range rangesToSplit {
			if rts == invalidValueRange {
				continue
			}
			below, eq, over := splitRange(rts, part.value)
			switch part.op {
			case "=":
				if eq != invalidValueRange {
					okRange := rg
					okRange.step = part.target
					okRange.setRanges(part.member, []valueRange{eq})
					rr = append(rr, okRange)
				}
				remainingRanges = append(remainingRanges, below, over)
			case "<":
				if below != invalidValueRange {
					okRange := rg
					okRange.step = part.target
					okRange.setRanges(part.member, []valueRange{below})
					rr = append(rr, okRange)
				}
				remainingRanges = append(remainingRanges, eq, over)
			case ">":
				if over != invalidValueRange {
					okRange := rg
					okRange.step = part.target
					okRange.setRanges(part.member, []valueRange{over})
					rr = append(rr, okRange)
				}
				remainingRanges = append(remainingRanges, below, eq)
			}
		}
		rg.setRanges(part.member, remainingRanges)
	}
	return rr
}

var (
	lightInput = rulesAndPieces{
		rules: []string{
			"px{a<2006:qkq,m>2090:A,rfg}",
			"pv{a>1716:R,A}",
			"lnx{m>1548:A,A}",
			"rfg{s<537:gd,x>2440:R,A}",
			"qs{s>3448:A,lnx}",
			"qkq{x<1416:A,crn}",
			"crn{x>2662:A,R}",
			"in{s<1351:px,qqz}",
			"qqz{s>2770:qs,m<1801:hdj,R}",
			"gd{a>3333:R,R}",
			"hdj{m>838:A,pv}",
		},
		pieces: []string{
			"{x=787,m=2655,a=1222,s=2876}",
			"{x=1679,m=44,a=2067,s=496}",
			"{x=2036,m=264,a=79,s=2244}",
			"{x=2461,m=1339,a=466,s=291}",
			"{x=2127,m=1623,a=2188,s=1013}",
		},
	}

	largeInput = rulesAndPieces{
		rules: []string{
			"hz{m>3518:A,A}",
			"xjq{s<700:R,x>3290:A,a>2004:R,R}",
			"dn{x<1908:R,a>539:A,s>1576:R,kdn}",
			"ql{m<3667:rpl,A}",
			"jsd{m>1643:R,R}",
			"dvq{s<1083:R,x>2321:A,A}",
			"qzq{x<3660:A,a>2909:jnb,vhm}",
			"pz{s>3001:jlf,zj}",
			"gb{s>236:tpj,mk}",
			"kgl{s<3549:nfm,a>2025:R,x>1769:A,A}",
			"jkc{s>1836:A,x>1826:A,vt}",
			"zn{a<471:A,s<757:lmg,th}",
			"vhm{s<3774:A,a<2437:A,A}",
			"gq{x>3007:ngs,xmr}",
			"lz{x>314:R,a>1284:R,A}",
			"gsg{m>1029:R,x>3382:R,R}",
			"cqj{x<2250:A,A}",
			"czh{x<2534:ntq,s>296:nvl,cq}",
			"sjb{m>910:A,s>2975:A,R}",
			"gss{m<3349:R,R}",
			"sf{m<1613:A,x>1749:R,s<2024:R,R}",
			"qd{m<2979:R,s>420:A,s<211:mmr,hc}",
			"nm{m<2530:A,a<690:R,A}",
			"cf{s<2923:A,A}",
			"gqh{m<690:gzc,ft}",
			"kxg{s<868:A,s>881:R,x<2889:R,R}",
			"bqk{s>1277:rqv,a>1262:xrn,xtd}",
			"jjf{s<2363:R,a>312:R,R}",
			"qlv{a>1656:rp,sx}",
			"czx{s<374:R,A}",
			"cvm{m<3623:A,a>3090:R,a<2883:R,A}",
			"np{x<57:R,R}",
			"df{a>2646:tct,qgd}",
			"pq{m>3886:A,s>1300:A,A}",
			"gvj{x<3344:A,m<3208:A,x<3652:R,A}",
			"ftk{s>183:czh,m<2812:dxs,mpj}",
			"gdl{a>2242:R,A}",
			"tpj{a>1952:R,x>414:A,a>1633:A,R}",
			"hd{m<746:A,x>3172:A,x<2654:lk,jjf}",
			"pqj{m<2823:rlq,m>3109:rj,A}",
			"ttj{s<1462:mj,A}",
			"vcl{x>2410:A,s>1991:A,x<2115:R,A}",
			"gsd{m>2665:A,x<2552:R,s<1074:A,A}",
			"dh{s<1213:A,bvg}",
			"xf{s>2328:pc,s<1137:jmc,tlf}",
			"jkd{x>1888:cf,x<908:kc,A}",
			"nxd{m>3569:A,x<552:R,A}",
			"kjv{x<2636:zt,x<3183:lkx,m>3114:qrb,pn}",
			"krp{s>3241:R,m<3498:A,A}",
			"smm{m<734:A,R}",
			"ph{a<3213:A,m>3216:R,A}",
			"nz{x<2386:dtq,s>2080:A,A}",
			"sgj{a<1835:R,A}",
			"qrb{a<955:cxr,m>3690:dc,tn}",
			"nk{x>795:qp,A}",
			"plm{m<1874:jbr,nkj}",
			"lc{x>454:A,x<266:R,A}",
			"jj{a>2841:A,a<2118:A,R}",
			"pj{m>3173:nxd,x<414:A,hdb}",
			"rh{x<3019:R,A}",
			"nfm{x>1631:R,A}",
			"blh{x<1029:pj,brl}",
			"ppt{a<3642:R,m<1223:gsg,A}",
			"hk{a<2029:hpn,s<492:zbb,jt}",
			"nmj{s<207:R,m<3830:A,x<2241:A,R}",
			"dl{a<871:A,m>2275:A,A}",
			"cbs{s<609:vx,s<657:R,A}",
			"vf{m>3366:A,a>560:R,a>296:A,A}",
			"xq{a>1033:R,a<956:A,A}",
			"pc{x>3394:A,x<3086:R,a>2727:A,A}",
			"nnk{m>2619:R,a>3006:hpp,jkh}",
			"zt{m>3191:zpb,bsq}",
			"vl{x<3532:A,s>887:A,R}",
			"lmg{m>3564:R,R}",
			"xrf{s>425:kq,m>2828:A,a>324:nm,fsk}",
			"tn{x<3644:zpz,m<3397:A,rch}",
			"hf{a<1446:A,x<2942:A,x<3022:A,R}",
			"bt{s<2790:lbf,a>1654:qpp,s>3238:nql,kjv}",
			"gg{a>2956:gsd,m<2713:dvq,R}",
			"mv{s>651:A,m<1448:A,x<146:A,A}",
			"tp{m>3253:gss,x>1334:A,ms}",
			"pg{x<2282:R,a<1349:vg,m>2840:R,hfl}",
			"rch{x>3834:A,s<3078:R,A}",
			"cxr{s>2975:plt,x>3547:R,x<3385:gr,fxd}",
			"bk{m<3130:hkl,m<3476:tp,m>3712:bgc,skz}",
			"xbx{a>3426:rc,s<863:sd,hq}",
			"kd{s>1340:A,A}",
			"bjf{a<593:R,A}",
			"mhm{a>1647:A,a>1415:A,m<2372:A,A}",
			"fx{a<3216:vz,srj}",
			"ml{a<2860:R,R}",
			"kzv{s<1301:A,R}",
			"jkh{s<947:R,R}",
			"pcl{m<3722:rvk,s>1134:pq,a<2815:R,A}",
			"bm{m<776:A,a>3600:R,A}",
			"hrt{s<674:R,x>2588:R,a>2018:R,A}",
			"lqj{x<676:kfp,x>856:lt,br}",
			"dhd{m<2607:R,A}",
			"qcv{x<1535:kl,R}",
			"szq{x>2379:hd,a>306:dn,s<1822:hkd,dns}",
			"tqq{a<522:mb,R}",
			"lkx{a<878:tqq,sbt}",
			"vcc{x>109:R,m<1781:R,m<1933:np,R}",
			"gdz{a>3532:A,A}",
			"vgt{m<1215:vcl,jcg}",
			"bxp{s>1485:A,s<563:R,m>1025:A,A}",
			"bmd{x<2839:rxh,a>2312:nmm,s>1414:krn,bq}",
			"rhg{s<152:A,R}",
			"cz{a<1432:ztv,xfr}",
			"vsd{a>2247:R,a>1946:R,a>1844:A,R}",
			"mqt{x<3172:R,x<3581:R,x>3825:R,R}",
			"tq{x<2593:R,A}",
			"fj{m<3481:rh,tv}",
			"gx{m>2883:R,R}",
			"ch{m<517:nk,pd}",
			"mbs{s>1219:A,s<1197:R,x>2230:R,R}",
			"rr{s<927:A,x>1452:A,x<1372:R,R}",
			"cgj{s<3063:R,A}",
			"bbv{s>1112:R,R}",
			"tqz{a>3803:nbb,psq}",
			"ncz{m>3717:R,a<634:R,R}",
			"jcg{s<2537:R,m<1343:R,A}",
			"jt{x<1325:pbb,R}",
			"kzb{m>3653:jqx,m>3482:gf,x>3249:xsk,rjx}",
			"rqv{m<3538:zzl,s<1338:tkq,R}",
			"hsb{a>1037:R,x>870:R,A}",
			"jb{x<2075:A,x<2506:lxt,A}",
			"vbk{a<2870:gjb,a>3277:A,dp}",
			"jqs{a<1324:A,m>2852:R,a>1529:R,A}",
			"qm{a>3573:A,R}",
			"rzk{x<557:vmp,s<83:tvm,m<3236:mc,ncz}",
			"ts{s<3733:fh,tfn}",
			"nsj{a<1336:A,a>1857:A,a>1584:A,R}",
			"vt{a>1102:R,m>2866:R,A}",
			"zkj{x>2500:R,a<467:A,R}",
			"bs{m<724:A,x>3306:fhb,x>3204:A,kdl}",
			"plt{x>3572:R,m<3702:R,s<3130:R,A}",
			"jfc{x>924:sf,m<1651:lz,mpg}",
			"knp{m>564:A,m<232:A,s>2339:A,R}",
			"qgd{a<1965:cm,A}",
			"tdm{m>2827:R,m>2718:A,a>3015:R,A}",
			"hgj{s>1445:zq,xrb}",
			"ztv{m>1357:qsl,tpx}",
			"fhb{s>1607:R,a<3027:R,A}",
			"khd{a>3834:R,m>3473:R,R}",
			"trc{s>512:rrk,m>2509:nd,s>326:sxh,cng}",
			"hr{a>656:A,m>2850:R,x>3389:A,A}",
			"rp{a>2022:R,R}",
			"gjb{a<2267:A,R}",
			"fd{x<1052:vtp,a<949:jrn,hk}",
			"fz{a>980:R,A}",
			"ccg{m>2905:R,a>2436:R,R}",
			"qlg{m>1037:R,R}",
			"nhh{x>2581:ql,m>3642:nmj,a>1170:kpf,jdt}",
			"jlf{x>3531:R,s>3219:A,A}",
			"nd{s<173:vbb,s>385:npn,bxz}",
			"fhj{m<724:R,a>3762:R,x>3053:R,R}",
			"tx{m>1717:A,m<1555:R,R}",
			"qz{x<3152:R,m<1947:R,tz}",
			"bln{m>1964:A,R}",
			"lh{s<2877:xqk,R}",
			"xhd{x<3255:ps,dz}",
			"krn{m>1099:gm,gjk}",
			"hpp{a<3169:A,s<938:R,A}",
			"srj{m<3640:A,A}",
			"fgh{m>1756:R,m>1722:R,s>1813:A,R}",
			"grd{x<1684:A,x>2484:A,x<2179:R,A}",
			"mpj{a<3142:A,x<2258:tpz,R}",
			"jv{x>3146:sq,s>269:R,cb}",
			"vc{a<1112:A,m<1598:jng,m<1694:A,fgh}",
			"tct{a>3518:R,jp}",
			"szk{x>2702:A,s>821:A,x<2020:A,R}",
			"tt{x<3002:A,x<3058:R,s<2845:R,A}",
			"xv{a<1832:A,s<1237:A,s<1251:A,R}",
			"vrq{a<1062:jtr,R}",
			"fbz{m>2837:A,A}",
			"ms{m<3212:A,A}",
			"dz{x>3696:dk,m<3179:A,A}",
			"sq{x<3517:A,m<1576:R,m>1616:A,R}",
			"vrv{x>2073:A,m<3335:A,a>3264:R,A}",
			"nx{s>284:R,A}",
			"bq{a<1964:ddh,zmr}",
			"ghk{a<441:R,A}",
			"br{a<2694:smm,x>795:zzv,s<2123:A,R}",
			"kmr{x>1717:A,a>3663:A,s<3538:R,R}",
			"hc{s<318:A,R}",
			"vr{a<399:A,x>2222:A,m>3016:A,A}",
			"cl{m>2530:ct,x<1855:qvv,dh}",
			"vmp{x>236:R,x>131:A,A}",
			"vq{s>526:vb,a>2478:kf,jv}",
			"mdb{a<380:R,m<2455:R,A}",
			"cpb{s<1179:dzj,dj}",
			"ssz{s>1128:A,R}",
			"mmr{m<3107:R,R}",
			"ddh{x>3498:tm,R}",
			"mk{x>561:A,s>98:R,x>197:R,R}",
			"jdn{a>1667:R,xlx}",
			"skh{s<2941:R,a<1171:R,s<3134:A,A}",
			"ktr{s>2122:rk,a<757:tsm,vrq}",
			"js{x<1491:A,R}",
			"jnr{s<796:R,s>1213:R,A}",
			"xd{x<2458:hv,x<3129:cgj,m>2807:qbb,pz}",
			"hfl{s<300:A,R}",
			"zxt{s<2782:A,s<3589:A,s>3804:R,R}",
			"kf{s<213:kht,s<412:nx,m>1587:A,ksp}",
			"jdt{a<424:R,x<2178:R,a<690:A,tvf}",
			"dns{s<2857:xbf,s<3399:llz,a>196:fdf,jsq}",
			"dd{x>1491:R,s<398:A,A}",
			"jmc{a<2988:xjq,s>688:A,trj}",
			"hkd{a<163:jnr,R}",
			"zmb{m<3165:pls,gpv}",
			"mb{m<3218:A,m>3585:A,m>3436:A,R}",
			"msx{x<3666:R,m<3398:R,a<3006:R,A}",
			"vz{s<1252:A,A}",
			"bvc{s<449:R,a>3035:A,R}",
			"gz{x<3316:zcz,qzq}",
			"bb{m<2381:R,x>2178:R,a<1435:R,A}",
			"qch{s>2944:A,m>1485:A,s>2282:R,R}",
			"fdf{a>234:R,a<219:R,m<652:R,R}",
			"bqr{m<1655:A,A}",
			"djn{x>2008:nt,a<993:jns,x<1955:R,sjr}",
			"xbf{m<712:A,x>1977:R,A}",
			"qpp{s<3361:xd,x>2233:gz,s>3593:ts,sxk}",
			"srq{s<567:nhh,s<733:kzb,a<1151:nfr,fj}",
			"vgz{x<2442:R,x>3163:A,s<1109:A,A}",
			"spf{m>3098:A,a>1114:R,x<2359:vr,R}",
			"lk{m<1007:R,m<1155:A,A}",
			"cq{x<3030:R,m<3304:A,s>244:R,R}",
			"hpn{m>3374:A,s<535:R,R}",
			"pnr{s<564:R,a>1040:A,x>573:knx,vd}",
			"cc{s>1430:mfd,A}",
			"fdb{s>349:R,s>228:A,s>81:A,A}",
			"cm{m<1985:R,x>2862:R,A}",
			"zg{x<2013:R,s>3441:A,R}",
			"th{s<768:A,s<776:R,m>3551:R,R}",
			"bvt{s<883:pxk,A}",
			"ps{x<2852:R,a<1334:kj,R}",
			"xmr{x>2490:xs,x<2141:djn,m>2911:spf,pg}",
			"tsm{s<872:A,a<426:bln,x>1500:db,A}",
			"qsl{m<1825:gd,m>1995:tmb,m>1914:ktr,plm}",
			"tvf{s<308:A,m<3477:A,x>2430:A,R}",
			"nbb{a>3887:R,m<2823:A,khd}",
			"jsq{m<806:R,a<125:R,R}",
			"gm{s<2956:R,m>1252:dfc,s>3550:R,R}",
			"kv{x>203:R,x<134:A,m<1920:R,R}",
			"bgc{m>3875:R,ssz}",
			"rxh{a<2551:vgt,jb}",
			"psq{s<967:dvk,m<3358:A,R}",
			"rg{x>249:R,x<137:A,A}",
			"rt{a>1268:R,a<1073:A,R}",
			"rcf{m<3157:A,x<3888:R,R}",
			"xtd{a<718:mbs,a>915:ss,a<801:R,A}",
			"tlf{m>331:A,a>2999:R,A}",
			"zl{a<2924:R,a<3347:A,m>592:R,R}",
			"kdl{x<3179:A,a>3060:A,A}",
			"xrn{s>1236:mdq,a<2044:A,gj}",
			"gj{a>2399:R,a<2219:A,s>1215:A,A}",
			"fq{s>518:A,x>3344:A,s>294:A,A}",
			"rb{a>3423:R,a>2952:A,x>915:R,R}",
			"bg{a<433:dd,A}",
			"zzv{x<828:R,s>2255:A,s>1152:R,R}",
			"lp{m<3249:A,s<3008:A,m<3278:A,R}",
			"cvd{a<2723:A,gdz}",
			"zcz{m>3081:A,a>2862:R,a>2373:A,R}",
			"qk{a<1277:R,m<3551:R,R}",
			"hj{m>3271:A,x<1007:A,R}",
			"tr{a<839:A,A}",
			"rld{s<823:R,m<2378:R,A}",
			"fzv{a<1192:mqd,a<1307:kg,nz}",
			"gpv{x>1141:R,jl}",
			"zbb{a>2244:ccg,x>1335:mmf,bh}",
			"ss{m<3642:A,A}",
			"zdd{m<2602:xt,tdm}",
			"zqq{m>575:R,A}",
			"zzz{a>3102:vrv,a<2800:cqj,a<2999:A,bvc}",
			"kfp{x>515:ml,A}",
			"dxf{a>315:R,R}",
			"vb{m>1613:jsd,a>2462:szk,hsz}",
			"lpk{s>1135:lht,s>989:gg,nnk}",
			"jtr{m<1957:A,a<935:R,A}",
			"ls{x>1788:zg,js}",
			"sd{s>376:lxx,ftk}",
			"ff{x>3648:R,a<3816:R,R}",
			"dvk{m>3204:A,R}",
			"nc{a<1279:A,x<2499:R,R}",
			"dtq{x<2009:A,R}",
			"rks{m<3222:R,R}",
			"in{m<2131:cz,mtv}",
			"jd{s>1478:td,m>1696:df,s<981:vq,vlr}",
			"ct{a>879:xv,m>3005:A,nv}",
			"lt{s>1942:zf,R}",
			"zfj{a<1034:zkj,rhg}",
			"dfc{m>1352:R,a>2004:R,x<3229:R,A}",
			"vrx{m<3175:xxm,R}",
			"ngs{a<867:lrj,x>3362:nj,qd}",
			"szg{a>2781:A,m<3156:R,a<2714:R,R}",
			"xg{m<2460:R,R}",
			"hh{x>2365:R,A}",
			"kg{m<614:tq,a<1245:R,x>2659:bxp,R}",
			"ht{s<3063:R,x<1179:A,R}",
			"gv{x<2021:R,R}",
			"mqd{m<718:A,s<2597:R,qxn}",
			"jng{m>1474:A,R}",
			"hzf{m>655:A,A}",
			"tdc{x<221:A,m>2804:R,A}",
			"tpx{x<1414:ch,a<816:szq,fzv}",
			"vxm{s<304:R,x<383:tdc,m<2773:mdb,A}",
			"dxs{x<2545:jcn,s<100:xg,s<153:A,jmj}",
			"gbb{x<2907:mvc,m<543:xf,x>3553:gqh,hjt}",
			"vg{s<356:A,m<2800:A,A}",
			"tkq{m>3809:A,s<1316:R,x>1546:R,R}",
			"zhp{m<1606:A,a>167:A,m<1712:A,R}",
			"zj{a<3026:A,s>2915:A,R}",
			"rkl{m>3409:R,A}",
			"rj{s>1336:A,R}",
			"skz{s<1142:vgz,m>3577:R,m<3533:nc,qk}",
			"bf{m>2781:zz,a>3752:A,a<3606:fdb,R}",
			"zs{a<2108:R,a<2506:R,tx}",
			"dj{s>1374:pzt,m>3269:bqk,s<1274:cl,sdn}",
			"dk{m>2966:R,s<971:R,m<2476:R,A}",
			"kdn{a<435:R,s<659:R,A}",
			"tv{a>1679:vsd,m<3686:R,x<2729:R,A}",
			"nfx{s<3034:A,x>3884:A,A}",
			"knx{x>816:R,x<702:A,R}",
			"rc{x<1712:fnm,tqz}",
			"jbr{a<852:A,R}",
			"cgb{a<789:A,s<1835:A,m>3623:rt,A}",
			"xqk{a<1905:R,a>2223:A,a>2093:A,A}",
			"rrk{s>719:rld,a<1094:cbs,qlv}",
			"bvg{m<2397:R,R}",
			"zb{x<3060:R,a<430:qr,csg}",
			"vd{x>278:A,R}",
			"qkc{a<3268:lzc,R}",
			"hsz{m<1538:A,s<813:A,A}",
			"fg{s>3518:R,a<3152:A,A}",
			"zf{a<2679:A,x>922:R,A}",
			"nfr{s<786:zn,s>828:qvz,zb}",
			"lsq{s<2128:R,x<1564:A,R}",
			"vlr{m>1602:kzv,s<1153:mqt,fvn}",
			"pzt{a<1316:hgj,x>2214:cc,ttj}",
			"fxd{a>391:R,a>191:R,s<2888:A,A}",
			"nt{s>446:R,x<2086:A,A}",
			"bts{a>638:R,A}",
			"lmt{a<2621:tt,a<3304:cx,fhj}",
			"ds{x<646:A,A}",
			"mpg{m<1752:A,A}",
			"tz{s<2648:A,a>3313:A,s<3179:A,A}",
			"brk{a<277:A,a>456:R,A}",
			"dc{m>3795:R,a>1285:cqc,m<3754:skh,A}",
			"rgq{a>637:A,x>2228:R,R}",
			"lht{a>3119:A,s>1330:A,R}",
			"mvc{a<2677:sg,fl}",
			"cng{x>2909:tc,s<113:kpq,s>224:fz,zfj}",
			"qbb{s<3159:hz,s<3283:krp,msx}",
			"lrj{x>3540:A,m<2947:fq,A}",
			"jcn{x>1192:A,m>2569:R,A}",
			"fcd{m<948:A,x<1316:ht,qch}",
			"qvz{m>3696:kxg,a<538:R,tr}",
			"mh{m<3314:grd,s<712:R,x<1541:A,cvm}",
			"gbz{a<622:R,A}",
			"ntq{s>250:A,m>3001:R,x<845:A,R}",
			"qr{s<813:R,R}",
			"rl{x>2198:A,x>2022:A,R}",
			"cs{a<815:A,s>2641:A,m<1444:A,R}",
			"jqx{a<941:A,a>1538:hrt,A}",
			"vx{m<2417:R,x>3122:R,A}",
			"gt{s>2144:bqr,a>368:A,x<2059:zhp,A}",
			"mdq{x<1604:R,x>2419:R,s<1260:R,A}",
			"gzc{m>604:hzf,zqq}",
			"kpq{x<2542:bb,x>2734:nxl,a>928:R,ghk}",
			"tvm{s>36:R,a<694:R,m<3038:R,R}",
			"cqc{m>3736:R,a>1421:A,A}",
			"cx{x>2986:R,A}",
			"nxl{m>2336:R,a>1436:R,R}",
			"lbf{a>1678:ks,mml}",
			"mx{m<1054:A,a<2830:R,A}",
			"xlx{m>2271:R,a>1109:R,x<3278:R,A}",
			"xs{x<2732:A,m>3031:czx,x<2854:gzj,R}",
			"km{x<2607:R,a>927:R,gvj}",
			"pf{x>1306:rr,mx}",
			"qj{m<2324:brk,R}",
			"jp{a>3169:R,s<612:R,a>2848:A,A}",
			"nmm{a<3403:sr,s<2153:sc,ppt}",
			"rpl{s>208:R,R}",
			"gjk{m>958:qlg,sjb}",
			"mc{s<140:R,s>158:A,a>577:R,A}",
			"qvv{s>1224:A,hsb}",
			"cb{m>1607:R,R}",
			"nvl{x<3257:R,a<3099:A,A}",
			"mmf{m<2967:A,A}",
			"kpf{s>198:rl,R}",
			"zmr{x<3388:vqn,R}",
			"rv{s<3061:R,R}",
			"nl{s<1931:pf,fcd}",
			"xrb{x<1432:R,A}",
			"kc{m<2063:R,A}",
			"rvk{s<1178:A,R}",
			"kfx{s>1262:R,x<1201:R,A}",
			"fb{s>2493:A,s>995:rg,m>1647:kv,mv}",
			"pls{x<1059:ds,s<3787:R,x<1733:A,fqj}",
			"sr{a>2831:R,s<1861:vl,R}",
			"bxz{s<250:dhd,A}",
			"nhx{m>3054:R,R}",
			"lhq{m>650:R,x>2996:zl,s<1699:R,R}",
			"zz{x>740:A,m>3478:A,s>359:R,R}",
			"zd{s<2863:knp,R}",
			"sl{a>3038:qvb,xj}",
			"bh{x<1198:R,m<2785:A,R}",
			"tpz{m<3568:R,R}",
			"vtp{s>401:pnr,a>1078:gb,s<174:rzk,vxm}",
			"sp{m<1518:cs,s<2529:R,m<1688:A,A}",
			"xx{x>3480:R,x>3304:R,R}",
			"gzj{m<2860:A,a>1265:R,A}",
			"fnm{s>574:hj,bf}",
			"ksp{s>456:A,R}",
			"pxk{a>3069:R,s>481:A,R}",
			"xxm{x>2290:R,A}",
			"csg{a>838:R,A}",
			"pn{m<2473:sqj,a<1054:lts,qg}",
			"bxt{m<3605:A,s>674:A,R}",
			"kvs{s<3401:km,rks}",
			"vqn{m<1234:R,s<899:A,m<1365:R,R}",
			"dzj{s>1078:bk,x>2290:xhd,blh}",
			"hgz{a<561:A,s<3039:R,x>1054:R,R}",
			"jns{s>319:A,m>2997:R,s<151:A,A}",
			"kj{m>3210:A,m>2599:R,A}",
			"jnb{x>3800:R,a>3516:A,R}",
			"pgg{x>1354:kd,m>2551:gbz,x>763:A,lc}",
			"sdn{a<991:pgg,s<1308:kzl,pqj}",
			"lxt{m<1151:R,x>2255:R,s>2561:A,R}",
			"nql{s<3500:kvs,x<2601:zmb,hp}",
			"nkj{x<1439:R,A}",
			"mfd{x<3094:A,x<3593:R,m<2791:R,R}",
			"npn{s<446:R,R}",
			"rk{a<925:A,A}",
			"fh{m<3144:R,a>2601:rb,R}",
			"kq{x>1327:R,s<628:R,s<719:R,R}",
			"llz{x<1865:R,x<2062:R,m<533:R,A}",
			"ft{s<1983:bbv,a<2323:gqm,a<3139:zxt,bm}",
			"mml{s>2364:vrx,m<3279:jkc,cgb}",
			"rlq{x>2132:R,m>2557:R,a>1746:A,A}",
			"prc{x>1012:nl,x>429:lqj,m>1338:nhr,dt}",
			"gqm{x>3726:A,s<2694:A,m<756:R,R}",
			"lvm{m<3367:A,s>598:A,x<2525:R,R}",
			"db{m>1949:R,x>2355:R,A}",
			"fvh{m<1036:R,A}",
			"fsj{m>1663:A,x>2977:A,s>2939:R,A}",
			"mtv{s>1535:bt,a>2595:xbx,s>888:cpb,zjb}",
			"zq{x>2291:R,R}",
			"fsk{s>265:R,m<2431:A,s>111:A,A}",
			"lts{x>3595:bjf,hr}",
			"tmb{s>1466:jkd,qcv}",
			"xfr{x<1561:prc,m<833:gbb,m>1471:jd,bmd}",
			"zjb{x<1868:fd,m>3209:srq,m>2729:gq,trc}",
			"gf{m<3555:R,x<2969:bxt,m<3599:A,R}",
			"hgd{x<3424:jqs,x<3784:hcv,rcf}",
			"gr{m>3624:A,a>594:A,a>338:R,A}",
			"hcv{a<1292:A,m>3015:R,R}",
			"nj{m>3046:A,x<3685:fbz,m>2877:R,A}",
			"rjx{s<635:lvm,A}",
			"jmj{m<2437:A,a<3023:R,A}",
			"hp{a>779:hgd,dxf}",
			"sqj{s<2972:R,m>2255:R,m<2212:rv,xx}",
			"vbb{s<93:A,a<1485:R,R}",
			"sc{x<3438:R,m<1135:A,a>3704:ff,qm}",
			"brl{s<979:A,a<1569:R,s>1040:R,A}",
			"hkl{a>1206:A,rgq}",
			"snt{a<3037:pcl,x>2181:qkc,fx}",
			"hv{m<2950:R,sdv}",
			"mj{a<1926:A,A}",
			"tm{s>687:R,A}",
			"sxh{x>2873:jdn,a>1033:mhm,a<613:qj,dl}",
			"sdv{m<3613:A,R}",
			"qg{s<2990:R,x<3664:gx,s<3085:nfx,R}",
			"vp{s<3458:R,a<3203:R,kmr}",
			"bsq{s<3005:R,A}",
			"qp{s<1699:A,A}",
			"jl{s>3798:A,R}",
			"fvn{x>2512:A,s<1276:R,R}",
			"zzl{m>3437:R,x>2355:R,R}",
			"qxn{x<3126:R,m<1095:A,R}",
			"pbb{m<3066:A,R}",
			"ks{m<3192:cvd,lsq}",
			"tfn{a>3183:nhx,A}",
			"trj{a<3362:A,m>290:A,x<3632:R,A}",
			"zpz{a<1354:R,s<3021:R,A}",
			"kl{m>2073:A,A}",
			"kht{a>3289:A,x<3178:R,A}",
			"xt{a>3063:A,x<2237:A,a>2887:A,A}",
			"fl{s>2668:fg,A}",
			"xsk{m<3311:nsj,a<1513:R,R}",
			"dt{s<2042:bvt,a>2530:zd,sgj}",
			"gd{a<620:gt,x>2444:vc,a>1108:jfc,sp}",
			"nv{s>1240:R,m>2836:R,a>353:A,R}",
			"td{a<2465:lh,m<1779:fsj,qz}",
			"hq{m<3025:lpk,m>3391:snt,sl}",
			"kzl{m>2636:A,m<2420:R,R}",
			"hdb{m<2705:A,m<2892:A,s<988:A,R}",
			"lzc{s<1244:R,A}",
			"pd{x>814:bts,fvh}",
			"fqj{a>577:R,x>2263:A,x>2013:R,A}",
			"hjt{x>3133:bs,s>2231:lmt,s>1243:lhq,jj}",
			"tc{a<1040:R,a<2061:R,R}",
			"sx{m>2430:R,a>1441:R,R}",
			"jrn{m<3177:xrf,bg}",
			"sjr{s>434:R,R}",
			"sbt{a>1333:hf,m<2974:xq,A}",
			"xj{s<1140:szg,a>2875:hh,m>3181:gv,A}",
			"zpb{m>3499:hgz,m<3296:lp,s>2976:vf,rkl}",
			"dp{a>3031:A,a>2953:A,R}",
			"lxx{s>575:mh,m<2984:zdd,zzz}",
			"sxk{x<1180:vbk,a>2579:vp,s<3501:ls,kgl}",
			"qvb{s<1124:R,x<2353:kfx,ph}",
			"sg{x>2195:R,m>320:gdl,a>2241:A,R}",
			"nhr{a>2999:fb,x<176:vcc,zs}",
		},
		pieces: []string{
			"{x=1853,m=1718,a=852,s=421}",
			"{x=1856,m=768,a=800,s=33}",
			"{x=2847,m=1317,a=3464,s=932}",
			"{x=2618,m=561,a=3141,s=132}",
			"{x=148,m=476,a=2620,s=457}",
			"{x=388,m=1384,a=860,s=100}",
			"{x=1929,m=115,a=349,s=290}",
			"{x=3086,m=2861,a=1622,s=48}",
			"{x=31,m=423,a=315,s=1698}",
			"{x=174,m=907,a=49,s=122}",
			"{x=541,m=15,a=242,s=2732}",
			"{x=1552,m=95,a=350,s=1981}",
			"{x=741,m=981,a=3076,s=2421}",
			"{x=849,m=166,a=1512,s=1803}",
			"{x=13,m=1454,a=146,s=2150}",
			"{x=957,m=67,a=56,s=360}",
			"{x=243,m=368,a=1375,s=878}",
			"{x=890,m=274,a=421,s=83}",
			"{x=474,m=87,a=2601,s=767}",
			"{x=993,m=106,a=3272,s=1520}",
			"{x=102,m=295,a=545,s=2670}",
			"{x=2084,m=1274,a=2583,s=1055}",
			"{x=1440,m=57,a=2201,s=1181}",
			"{x=189,m=4,a=515,s=3434}",
			"{x=164,m=15,a=38,s=368}",
			"{x=643,m=2265,a=1169,s=1196}",
			"{x=133,m=323,a=36,s=737}",
			"{x=1924,m=859,a=590,s=268}",
			"{x=1308,m=668,a=627,s=64}",
			"{x=1277,m=1203,a=2822,s=164}",
			"{x=143,m=1445,a=1323,s=1941}",
			"{x=876,m=577,a=159,s=2538}",
			"{x=877,m=664,a=121,s=238}",
			"{x=578,m=1677,a=99,s=825}",
			"{x=94,m=697,a=629,s=786}",
			"{x=1108,m=1064,a=5,s=597}",
			"{x=416,m=2871,a=946,s=208}",
			"{x=1055,m=20,a=1258,s=1102}",
			"{x=85,m=144,a=1934,s=120}",
			"{x=747,m=2995,a=841,s=809}",
			"{x=205,m=94,a=959,s=1002}",
			"{x=251,m=1836,a=475,s=381}",
			"{x=363,m=765,a=187,s=536}",
			"{x=1823,m=1509,a=361,s=1068}",
			"{x=136,m=765,a=260,s=899}",
			"{x=1752,m=178,a=310,s=227}",
			"{x=318,m=617,a=1396,s=564}",
			"{x=1371,m=196,a=2487,s=2149}",
			"{x=2723,m=926,a=1502,s=1746}",
			"{x=146,m=448,a=181,s=2032}",
			"{x=1501,m=2536,a=1073,s=476}",
			"{x=850,m=900,a=29,s=2148}",
			"{x=984,m=128,a=1750,s=1273}",
			"{x=65,m=740,a=648,s=1147}",
			"{x=660,m=1068,a=1249,s=1061}",
			"{x=289,m=1612,a=710,s=1181}",
			"{x=1872,m=258,a=1788,s=508}",
			"{x=492,m=1178,a=125,s=618}",
			"{x=849,m=1586,a=3172,s=1776}",
			"{x=398,m=1499,a=737,s=645}",
			"{x=115,m=227,a=154,s=622}",
			"{x=415,m=505,a=388,s=338}",
			"{x=104,m=1398,a=921,s=164}",
			"{x=421,m=1201,a=3389,s=456}",
			"{x=661,m=237,a=292,s=283}",
			"{x=143,m=845,a=1391,s=1900}",
			"{x=777,m=141,a=207,s=3834}",
			"{x=780,m=927,a=62,s=430}",
			"{x=2947,m=1361,a=5,s=50}",
			"{x=3493,m=134,a=1163,s=2043}",
			"{x=135,m=456,a=58,s=1093}",
			"{x=1244,m=758,a=450,s=1089}",
			"{x=290,m=2583,a=1692,s=1164}",
			"{x=200,m=713,a=192,s=2113}",
			"{x=1164,m=57,a=3464,s=2020}",
			"{x=1023,m=137,a=1328,s=1460}",
			"{x=612,m=597,a=101,s=427}",
			"{x=2247,m=891,a=1224,s=2817}",
			"{x=870,m=3099,a=1275,s=305}",
			"{x=2099,m=1353,a=1867,s=55}",
			"{x=264,m=153,a=2560,s=1307}",
			"{x=274,m=125,a=167,s=27}",
			"{x=208,m=88,a=1676,s=1450}",
			"{x=5,m=161,a=3312,s=1403}",
			"{x=2294,m=1021,a=4,s=1766}",
			"{x=2683,m=91,a=441,s=2454}",
			"{x=503,m=1775,a=492,s=2809}",
			"{x=517,m=19,a=609,s=1051}",
			"{x=603,m=313,a=211,s=2889}",
			"{x=1663,m=342,a=1651,s=1501}",
			"{x=305,m=1190,a=232,s=1049}",
			"{x=3469,m=86,a=883,s=1897}",
			"{x=654,m=1526,a=741,s=186}",
			"{x=2611,m=959,a=541,s=1131}",
			"{x=3122,m=827,a=1701,s=1953}",
			"{x=594,m=1973,a=475,s=191}",
			"{x=134,m=1169,a=125,s=554}",
			"{x=4,m=401,a=176,s=3228}",
			"{x=2978,m=7,a=309,s=3088}",
			"{x=1187,m=2241,a=221,s=185}",
			"{x=666,m=1214,a=144,s=197}",
			"{x=176,m=30,a=208,s=3302}",
			"{x=1581,m=2530,a=596,s=244}",
			"{x=1264,m=1345,a=16,s=53}",
			"{x=122,m=424,a=2194,s=3612}",
			"{x=859,m=7,a=8,s=2258}",
			"{x=1346,m=211,a=113,s=34}",
			"{x=663,m=1448,a=2323,s=1344}",
			"{x=112,m=141,a=708,s=2131}",
			"{x=312,m=1345,a=1836,s=337}",
			"{x=5,m=1556,a=1244,s=848}",
			"{x=1589,m=233,a=21,s=933}",
			"{x=2173,m=2390,a=180,s=864}",
			"{x=1645,m=1798,a=773,s=297}",
			"{x=3292,m=109,a=1124,s=613}",
			"{x=482,m=1353,a=784,s=3347}",
			"{x=1189,m=3164,a=1874,s=1053}",
			"{x=495,m=431,a=831,s=35}",
			"{x=686,m=915,a=1823,s=809}",
			"{x=766,m=1004,a=1354,s=307}",
			"{x=3352,m=1047,a=471,s=148}",
			"{x=155,m=1449,a=997,s=345}",
			"{x=117,m=93,a=1355,s=14}",
			"{x=1710,m=1171,a=875,s=1402}",
			"{x=1339,m=1068,a=2676,s=354}",
			"{x=1306,m=29,a=1186,s=2010}",
			"{x=179,m=532,a=581,s=1137}",
			"{x=349,m=2778,a=1035,s=1522}",
			"{x=1779,m=831,a=91,s=447}",
			"{x=2267,m=267,a=370,s=177}",
			"{x=684,m=3595,a=349,s=55}",
			"{x=3394,m=422,a=1182,s=1468}",
			"{x=1902,m=359,a=956,s=2143}",
			"{x=3729,m=1383,a=799,s=887}",
			"{x=2182,m=855,a=1277,s=195}",
			"{x=608,m=1985,a=83,s=3923}",
			"{x=2768,m=233,a=1538,s=2232}",
			"{x=59,m=1225,a=270,s=983}",
			"{x=1804,m=2039,a=957,s=705}",
			"{x=460,m=1115,a=1049,s=376}",
			"{x=27,m=144,a=1421,s=2553}",
			"{x=838,m=2998,a=563,s=3050}",
			"{x=142,m=466,a=479,s=1125}",
			"{x=715,m=565,a=32,s=138}",
			"{x=156,m=668,a=507,s=2073}",
			"{x=1926,m=267,a=3414,s=554}",
			"{x=1253,m=1693,a=2019,s=360}",
			"{x=398,m=2354,a=1778,s=643}",
			"{x=954,m=525,a=1508,s=1607}",
			"{x=37,m=31,a=149,s=803}",
			"{x=212,m=383,a=1288,s=145}",
			"{x=39,m=639,a=364,s=1277}",
			"{x=2519,m=1272,a=31,s=2869}",
			"{x=162,m=1170,a=449,s=516}",
			"{x=2101,m=1142,a=2348,s=156}",
			"{x=915,m=2517,a=1793,s=2079}",
			"{x=540,m=856,a=2704,s=3}",
			"{x=414,m=836,a=71,s=1790}",
			"{x=2796,m=1429,a=60,s=928}",
			"{x=52,m=88,a=610,s=517}",
			"{x=1019,m=1532,a=2767,s=632}",
			"{x=1443,m=441,a=228,s=642}",
			"{x=328,m=595,a=947,s=936}",
			"{x=463,m=2650,a=104,s=3234}",
			"{x=789,m=62,a=281,s=257}",
			"{x=293,m=18,a=56,s=62}",
			"{x=9,m=665,a=556,s=426}",
			"{x=1141,m=1351,a=760,s=599}",
			"{x=370,m=839,a=102,s=1985}",
			"{x=2295,m=197,a=2136,s=181}",
			"{x=545,m=1190,a=1118,s=1373}",
			"{x=1488,m=1075,a=265,s=1678}",
			"{x=890,m=3,a=376,s=406}",
			"{x=75,m=318,a=14,s=264}",
			"{x=920,m=63,a=238,s=3394}",
			"{x=1787,m=3530,a=2762,s=5}",
			"{x=316,m=1158,a=1934,s=1069}",
			"{x=573,m=195,a=105,s=564}",
			"{x=1821,m=2141,a=579,s=808}",
			"{x=323,m=2219,a=61,s=208}",
			"{x=140,m=1375,a=46,s=2408}",
			"{x=358,m=529,a=220,s=31}",
			"{x=203,m=789,a=585,s=868}",
			"{x=2118,m=884,a=828,s=362}",
			"{x=388,m=2794,a=2062,s=372}",
			"{x=862,m=2303,a=1032,s=196}",
			"{x=725,m=1153,a=478,s=1423}",
			"{x=1353,m=313,a=2826,s=31}",
			"{x=2985,m=183,a=1256,s=734}",
			"{x=446,m=417,a=1970,s=122}",
			"{x=3598,m=2237,a=38,s=247}",
			"{x=455,m=1138,a=109,s=527}",
			"{x=697,m=1815,a=3009,s=530}",
			"{x=30,m=49,a=497,s=871}",
			"{x=866,m=400,a=1041,s=2446}",
			"{x=304,m=512,a=1530,s=194}",
			"{x=1213,m=2841,a=152,s=553}",
			"{x=477,m=547,a=534,s=2815}",
			"{x=411,m=765,a=70,s=1005}",
			"{x=1767,m=1973,a=872,s=494}",
		},
	}
)

func compare(value int, op string, testValue int) bool {
	switch op {
	case "=":
		return value == testValue
	case ">":
		return value > testValue
	case "<":
		return value < testValue
	default:
		panic("unsupported")
	}
}

func parseInput(input rulesAndPieces) (map[string]rule, []piece) {
	pieces := make([]piece, len(input.pieces))
	for i, p := range input.pieces {
		parts := strings.Split(p[1:len(p)-1], ",")
		var pc piece
		for _, part := range parts {
			var err error
			switch part[0] {
			case 'x':
				pc.x, err = strconv.Atoi(part[2:])
			case 'm':
				pc.m, err = strconv.Atoi(part[2:])
			case 'a':
				pc.a, err = strconv.Atoi(part[2:])
			case 's':
				pc.s, err = strconv.Atoi(part[2:])
			}
			if err != nil {
				panic(err)
			}
		}
		pieces[i] = pc
	}

	rules := make(map[string]rule)
	for _, r := range input.rules {
		name := r[:strings.IndexByte(r, '{')]
		parts := strings.Split(r[len(name)+1:len(r)-1], ",")
		ruleItems := make([]rulePart, len(parts))
		for i, p := range parts {
			var err error
			if condAndTarget := strings.Split(p, ":"); len(condAndTarget) == 2 {
				ruleItems[i] = rulePart{
					member: condAndTarget[0][0:1],
					op:     condAndTarget[0][1:2],
					target: condAndTarget[1],
				}
				ruleItems[i].value, err = strconv.Atoi(condAndTarget[0][2:])
				if err != nil {
					panic(err)
				}
			} else {
				ruleItems[i] = rulePart{
					target: condAndTarget[0],
				}
			}
		}
		rules[name] = rule{
			apply: func(p piece) string {
				for _, item := range ruleItems {
					switch item.member {
					case "x":
						if compare(p.x, item.op, item.value) {
							return item.target
						}
					case "m":
						if compare(p.m, item.op, item.value) {
							return item.target
						}
					case "a":
						if compare(p.a, item.op, item.value) {
							return item.target
						}
					case "s":
						if compare(p.s, item.op, item.value) {
							return item.target
						}
					case "":
						return item.target
					default:
						panic("unknown")
					}
				}
				panic("invalid entry")
			},
			parts: ruleItems,
		}
	}

	return rules, pieces
}

func computePiecesIndex(input rulesAndPieces) {
	rules, pieces := parseInput(input)

	value := 0
	for _, p := range pieces {
		step := "in"
		for step != "A" && step != "R" {
			step = rules[step].apply(p)
		}
		if step == "A" {
			value += p.x + p.m + p.a + p.s
		}
	}
	fmt.Printf("Score is : %d\n", value)
}

type valueRange struct {
	min int
	max int
}

var invalidValueRange valueRange

type ruleRange struct {
	step    string
	xRanges []valueRange
	mRanges []valueRange
	aRanges []valueRange
	sRanges []valueRange
}

func (rr ruleRange) ranges(member string) []valueRange {
	switch member {
	case "x":
		return rr.xRanges
	case "m":
		return rr.mRanges
	case "a":
		return rr.aRanges
	case "s":
		return rr.sRanges
	}
	panic("invalid member")
}
func (rr *ruleRange) setRanges(member string, rgs []valueRange) {
	switch member {
	case "x":
		rr.xRanges = rgs
	case "m":
		rr.mRanges = rgs
	case "a":
		rr.aRanges = rgs
	case "s":
		rr.sRanges = rgs
	default:
		panic("invalid member")
	}
}

func (rr ruleRange) Count() int {
	pcsCount := 1
	for _, ranges := range [][]valueRange{rr.xRanges, rr.mRanges, rr.aRanges, rr.sRanges} {
		rangeSize := 0
		for _, rg := range ranges {
			if rg == invalidValueRange {
				continue
			}
			if rg.max >= rg.min {
				rangeSize += rg.max - rg.min + 1
			}
		}
		pcsCount *= rangeSize
	}
	return pcsCount
}

func computePiecesRange(input rulesAndPieces) int {
	rules, _ := parseInput(input)

	const maxRange = 4000
	ranges := []ruleRange{
		{
			step:    "in",
			xRanges: []valueRange{{1, maxRange}},
			mRanges: []valueRange{{1, maxRange}},
			aRanges: []valueRange{{1, maxRange}},
			sRanges: []valueRange{{1, maxRange}},
		},
	}

	value := 0
	valueRejected := 0
	for i := 0; i < len(ranges); i++ {
		r := ranges[i]
		if r.step == "R" || r.step == "A" {
			pcsCount := r.Count()
			if r.step == "A" {
				value += pcsCount
			} else {
				valueRejected += pcsCount
			}
			continue
		}

		split := rules[r.step].ranges(r)
		count := r.Count()
		for _, sr := range split {
			sc := sr.Count()
			count -= sc
		}
		if count != 0 {
			fmt.Println("break here")
			_ = rules[r.step].ranges(r)
		}
		ranges = append(ranges, split...)
	}

	if value+valueRejected != maxRange*maxRange*maxRange*maxRange {
		fmt.Printf("inconsistent sorting")
	}
	fmt.Printf("Num pieces is : %d\n", value)
	return value
}
