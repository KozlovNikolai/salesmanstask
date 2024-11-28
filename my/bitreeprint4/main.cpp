#include <string>
#include <iostream>
#include <vector>
#include <functional>
#include <algorithm>
#include <tuple>
#include <stdexcept>
#include <memory>
#ifdef _WIN32
    #include <windows.h>
#else
    #include <locale.h>
#endif

struct BinTree {
    int value;
    BinTree* left;
    BinTree* right;
};

#define LN { std::cout << __LINE__ << std::endl; }
#define DEB(x) { std::cout << #x << "=(" << (x) << ") "; }

// https://en.wikipedia.org/wiki/Box_Drawing_(Unicode_block)
// static std::string ch_hor = "-", ch_ver = "|", ch_ddia = "/", ch_rddia = "\\", ch_udia = "\\", ch_ver_hor = "|-", ch_udia_hor = "\\-", ch_ddia_hor = "/-", ch_ver_spa = "| ";
static std::string ch_hor = "\u2500" /*─*/, ch_ver = "\u2502" /*│*/, ch_ddia = "\u250C" /*┌*/, ch_rddia = "\u2510" /*┐*/, ch_udia = "\u2514" /*└*/, ch_ver_hor = "\u251C\u2500" /*├─*/, ch_udia_hor = "\u2514\u2500" /*└─*/, ch_ddia_hor = "\u250C\u2500" /*┌─*/, ch_ver_spa = "\u2502 " /*│ */;

void dump0(BinTree const * node, std::string const & prefix = "", bool root = true, bool last = true) {
    std::cout << prefix << (root ? "" : (last ? ch_udia_hor : ch_ver_hor)) << (node ? std::to_string(node->value) : "") << std::endl;
    if (!node || (!node->left && !node->right))
        return;
    std::vector<BinTree*> v{node->left, node->right};
    for (size_t i = 0; i < v.size(); ++i)
        dump0(v[i], prefix + (root ? "" : (last ? "  " : ch_ver_spa)), false, i + 1 >= v.size());
}

void dump2(BinTree const * node, std::string const & rpref = "", std::string const & cpref = "", std::string const & lpref = "") {
    if (!node) return;
    if (node->right)
        dump2(node->right, rpref + "  ", rpref + ch_ddia_hor, rpref + ch_ver_spa);
    std::cout << cpref << std::to_string(node->value) << std::endl;
    if (node->left)
        dump2(node->left, lpref + ch_ver_spa, lpref + ch_udia_hor, lpref + "  ");
}

void dump4(BinTree const * node, bool high = true, std::vector<std::string> const & lpref = std::vector<std::string>(), std::vector<std::string> const & cpref = std::vector<std::string>(), std::vector<std::string> const & rpref = std::vector<std::string>(), bool root = true, bool left = true, std::shared_ptr<std::vector<std::vector<std::string>>> lines = nullptr) {
    if (!node) return;
    typedef std::vector<std::string> VS;
    auto VSCat = [](VS const & a, VS const & b){ auto r = a; r.insert(r.end(), b.begin(), b.end()); return r; };
    if (root) lines = std::make_shared<std::vector<VS>>();
    if (node->left)
        dump4(node->left, high, VSCat(lpref, high ? VS({" ", " "}) : VS({" "})), VSCat(lpref, high ? VS({ch_ddia, ch_ver}) : VS({ch_ddia})), VSCat(lpref, high ? VS({ch_hor, " "}) : VS({ch_hor})), false, true, lines);
    auto sval = std::to_string(node->value);
    size_t sm = left || sval.empty() ? sval.size() / 2 : ((sval.size() + 1) / 2 - 1);
    for (size_t i = 0; i < sval.size(); ++i)
        lines->push_back(VSCat(i < sm ? lpref : i == sm ? cpref : rpref, {std::string(1, sval[i])}));
    if (node->right)
        dump4(node->right, high, VSCat(rpref, high ? VS({ch_hor, " "}) : VS({ch_hor})), VSCat(rpref, high ? VS({ch_rddia, ch_ver}) : VS({ch_rddia})), VSCat(rpref, high ? VS({" ", " "}) : VS({" "})), false, false, lines);
    if (root) {
        VS out;
        for (size_t l = 0;;++l) {
            bool last = true;
            std::string line;
            for (size_t i = 0; i < lines->size(); ++i)
                if (l < (*lines).at(i).size()) {
                    line += lines->at(i)[l];
                    last = false;
                } else line += " ";
            if (last) break;
            out.push_back(line);
        }
        for (size_t i = 0; i < out.size(); ++i)
            std::cout << out[i] << std::endl;
    }
}

void dump3(BinTree * root, int space = 0) {
   if (!root)
      return;
   enum { COUNT = 2 };
   space += COUNT;
   dump3(root->right, space);
   for (int i = COUNT; i < space; ++i)
      std::cout << "  ";
   std::cout << root->value << std::endl;
   dump3(root->left, space);
}

void dump1(BinTree const * node) {
    #define _MAX(x, y) ((x) > (y) ? (x) : (y))
    #define _MIN(x, y) ((x) < (y) ? (x) : (y))

    auto RepStr = [](std::string const & s, size_t cnt) {
        if (ptrdiff_t(cnt) < 0)
            throw std::runtime_error("RepStr: Bad value " + std::to_string(ptrdiff_t(cnt)) + "!");
        std::string r;
        for (size_t i = 0; i < cnt; ++i)
            r += s;
        return r;
    };
    std::function<std::tuple<std::vector<std::string>, size_t, size_t>(BinTree const * node, bool)> Rec;
    Rec = [&RepStr, &Rec](BinTree const * node, bool left){
        std::vector<std::string> lines;
        if (!node)
            return std::make_tuple(lines, size_t(0), size_t(0));
        auto sval = std::to_string(node->value);
        //if (sval.size() % 2 != 1) sval += " ";
        auto resl = Rec(node->left, true), resr = Rec(node->right, false);
        auto const & vl = std::get<0>(resl);
        auto const & vr = std::get<0>(resr);
        auto cl = std::get<1>(resl), cr = std::get<1>(resr), lss = std::get<2>(resl), rss = std::get<2>(resr);
        size_t lv = sval.size();
        size_t ls = vl.size() > 0 ? lss : size_t(0);
        size_t rs = vr.size() > 0 ? rss : size_t(0);
        size_t lis = ls == 0 ? lv / 2 : _MAX(int(lv / 2 + 1 - (ls - cl)), 0);
        size_t ris = rs == 0 ? (lv + 1) / 2 : _MAX(int((lv + 1) / 2 - cr), (lis > 0 ? 0 : 1));
        size_t dashls = (ls == 0 ? 0 : ls - cl - 1 + lis - lv / 2), dashrs = (rs == 0 ? 0 : cr + ris - (lv + 1) / 2);
        //DEB(node->value); DEB(lv); DEB(ls); DEB(rs); DEB(cl); DEB(cr); DEB(lis); DEB(ris); DEB(dashls); DEB(dashrs); std::cout << std::endl;
        lines.push_back(
            (ls == 0 ? "" : (RepStr(" ", cl) + ch_ddia + RepStr(ch_hor, dashls))) +
            sval + (rs == 0 ? "" : (RepStr(ch_hor, dashrs) + ch_rddia + RepStr(" ", rs - cr - 1)))
        );
        //if (lines.back().size() != ls + lis + ris + rs)
        //    throw std::runtime_error("Dump: First line wrong size, got " + std::to_string(lines.back().size()) + ", expected " + std::to_string(ls + lis + ris + rs));
        for (size_t i = 0; i < _MAX(vl.size(), vr.size()); ++i) {
            std::string sl = RepStr(" ", ls), sr = RepStr(" ", rs);
            if (i < vl.size()) sl = vl[i];
            if (i < vr.size()) sr = vr[i];
            sl = sl + RepStr(" ", lis);
            sr = RepStr(" ", ris) + sr;
            lines.push_back(sl + sr);
        }
        return std::make_tuple(lines, (left || ls + lis == 0 || lv % 2 == 1) ? ls + lis : ls + lis - 1, ls + lis + ris + rs);
    };
    auto v = std::get<0>(Rec(node, true));
    for (size_t i = 0; i < v.size(); ++i)
        std::cout << v[i] << std::endl;
        
    #undef _MAX
    #undef _MIN
}

int main() {
    #ifdef _WIN32
        SetConsoleOutputCP(65001);
    #else
        setlocale(LC_ALL, "en_US.UTF-8");
    #endif
    try {
        auto tree = 
            new BinTree{10,
                new BinTree{5,
                    new BinTree{1, 0,
                        new BinTree{2, 0, 0}
                    },
                    new BinTree{6, 0,
                        new BinTree{8, new BinTree{7, 0, 0}, 0},
                    },
                },
                new BinTree{19,
                    new BinTree{17, 0, 0},
                    new BinTree{21, new BinTree{20, 0, 0}, new BinTree{250, 0, 0}},
                },
            };
        std::cout << "===dump0===" << std::endl;
        dump0(tree);
        std::cout << "===dump1===" << std::endl;
        dump1(tree);
        std::cout << "===dump2===" << std::endl;
        dump2(tree);
        std::cout << "===dump3===" << std::endl;
        dump3(tree);
        std::cout << "===dump4_high===" << std::endl;
        dump4(tree, true);
        std::cout << "===dump4_low===" << std::endl;
        dump4(tree, false);
        return 0;
    } catch (std::exception const & ex) {
        std::cout << "Exception: " << ex.what() << std::endl;
        return -1;
    }
}