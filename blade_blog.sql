SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- 数据库: `blade_blog`
--

-- --------------------------------------------------------

--
-- 表的结构 `article`
--

CREATE TABLE IF NOT EXISTS `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `content` text NOT NULL,
  `alias` varchar(200) DEFAULT NULL,
  `created` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=32 ;

--
-- 转存表中的数据 `article`
--

INSERT INTO `article` (`id`, `title`, `content`, `alias`, `created`) VALUES
(1, '这是标题1', '内容1', 'alias1', 1555685803),
(2, '这是标题2', '内容2', 'alias2', 1555597025),
(3, '这是标题3', '内容3', 'alias3', 1555597045),
(4, '这是标题4', '内容4', 'alias4', 1555597082),
(5, '这是标题5', '内容5', 'alias5', 1555597090),
(6, '这是标题6', '内容6', 'alias6', 1555597096),
(7, '这是标题7', '内容7', 'alias7', 1555597102),
(8, '这是标题8', '内容8', 'alias8', 1555597108),
(9, '这是标题9', '内容9', 'alias9', 1555597114),
(10, '这是标题10', '内容10', 'alias10', 1555597122),
(11, '这是标题11', '内容11', 'alias11', 1555597130),
(12, '这是标题12', '内容12', 'alias12', 1555597136),
(13, '这是标题13', '内容13', 'alias13', 1555597144),
(14, '这是标题14', '内容14', 'alias14', 1555597151),
(15, '这是标题15', '内容15', 'alias15', 1555597157),
(16, '这是标题16', '内容16', 'alias16', 1555597162),
(17, '这是标题17', '内容17', 'alias17', 1555597169),
(18, '这是标题18', '内容18', 'alias18', 1555597175),
(19, '这是标题19', '内容19', 'alias19', 1555597182),
(20, '这是标题20', '内容20', 'alias20', 1555597192),
(21, 'Java浅克隆与深克隆', '## 浅克隆\n> 浅克隆只复制当前对象的所有基本数据类型，以及相应的引用变量，但没有复制引用变量指向的实际对象，也就是只复制了引用变量的内存地址。\n\n重写Object的clone方法，然后实现Cloneable接口。被克隆的类必须实现Cloneable接口，否则如果我们将在对象上调用clone()时，JVM将抛出CloneNotSupportedException。\n\n### Main.java\n```java\npublic class Main {\n    public static void main(String[] args) {\n        Person source = new Person();\n        Person target = null;\n        try {\n            target = source.clone();\n        } catch (CloneNotSupportedException e) {\n            e.printStackTrace();\n        }\n        System.out.println("被克隆对象: " + source.hashCode() + "\n" +\n                "被克隆对象内的dog属性: " + source.getDog().hashCode());\n        System.out.println("===============================");\n        System.out.println("克隆出来的对象: " + target.hashCode() + "\n" +\n                "克隆出来的对象内的dog属性: " + target.getDog().hashCode());\n    }\n}\n```\n### Person.java\n```java\npublic class Person implements Cloneable {\n    private Dog dog;\n\n    public Dog getDog() {\n        return dog;\n    }\n\n    public Person() {\n        this.dog = new Dog();\n    }\n\n    @Override\n    protected Person clone() throws CloneNotSupportedException {\n        return (Person) super.clone();\n    }\n}\n```\n### 输出\n```\n被克隆对象: 1163157884\n被克隆对象内的dog属性: 1956725890\n===============================\n克隆出来的对象: 356573597\n克隆出来的对象内的dog属性: 1956725890\n```\n我们发现两个对象的dog属性在内存中的标识并没有改变，因为浅克隆只拷贝了引用变量所指向堆内存的内存地址。\n\n## 深克隆\n> 深克隆彻底复制了当前对象，此对象与母对象在任何引用路径上都不存在共享的实例对象。也就是说深克隆把所有引用变量所指向的变量都拷贝了一份。\n\n先把对象序列化到流中，然后再把对象从流中取出来，取出来的对象和原来的对象就没有联系了，引用变量所指向的地址也和原来的对象不同，这样就达到了深克隆的目的。不过这种方式的效率会比较低。\n除了通过序列化成流进行深克隆，还可以通过手动设置属性的值实现克隆。\n下面演示通过序列化成流进行深克隆。\n\n### Main.java\n```java\npublic class Main {\n    public static void main(String[] args) {\n        Person source = new Person();\n        Person target = (Person) CloneUtils.clone(source);\n        System.out.println("被克隆对象: " + source.hashCode() + "\n" +\n                "被克隆对象内的dog属性: " + source.getDog().hashCode());\n        System.out.println("===============================");\n        System.out.println("克隆出来的对象: " + target.hashCode() + "\n" +\n                "克隆出来的对象内的dog属性: " + target.getDog().hashCode());\n    }\n}\n```\n### Person.java\n```java\npublic class Person implements Serializable {\n    private Dog dog;\n\n    public Dog getDog() {\n        return dog;\n    }\n\n    public Person() {\n        this.dog = new Dog();\n    }\n}\n```\n### Dog.java\n```java\npublic class Dog implements Serializable {\n    private String name;\n\n    public Dog() {\n        this.name = "大黄";\n    }\n}\n```\n### CloneUtils.java\n```java\npublic class CloneUtils {\n\n    public static Object clone(Object obj) {\n        Object cloneObj = null;\n        try {\n            //写入字节流\n            ByteArrayOutputStream out = new ByteArrayOutputStream();\n            ObjectOutputStream obs = new ObjectOutputStream(out);\n            obs.writeObject(obj);\n            obs.close();\n\n            //分配内存，写入原始对象，生成新对象\n            ByteArrayInputStream ios = new ByteArrayInputStream(out.toByteArray());\n            ObjectInputStream ois = new ObjectInputStream(ios);\n\n            //返回生成的新对象\n            cloneObj = ois.readObject();\n            ois.close();\n\n        } catch (Exception e) {\n            e.printStackTrace();\n        }\n        return cloneObj;\n    }\n}\n```\n### 输出\n```\n被克隆对象: 1836019240\n被克隆对象内的dog属性: 325040804\n===============================\n克隆出来的对象: 363771819\n克隆出来的对象内的dog属性: 2065951873\n```\n由输出结果可以看出引用变量dog的值已经指向了另外的一个堆内存地址，即通过序列化成流实现了深克隆。', 'java-clone', 1555766309);

-- --------------------------------------------------------

--
-- 表的结构 `config`
--

CREATE TABLE IF NOT EXISTS `config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) DEFAULT NULL,
  `value` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=5 ;

--
-- 转存表中的数据 `config`
--

INSERT INTO `config` (`id`, `name`, `value`) VALUES
(1, 'title', '我的博客'),
(2, 'keywords', 'blade,blog,java');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2 ;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `username`, `password`) VALUES
(1, 'admin', '21232f297a57a5a743894a0e4a801fc3');

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
