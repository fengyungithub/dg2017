<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: articles.proto

namespace Articles;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>articles.ListResponse</code>
 */
class ListResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>repeated .articles.Article articles = 1;</code>
     */
    private $articles;

    public function __construct() {
        \GPBMetadata\Articles::initOnce();
        parent::__construct();
    }

    /**
     * <code>repeated .articles.Article articles = 1;</code>
     */
    public function getArticles()
    {
        return $this->articles;
    }

    /**
     * <code>repeated .articles.Article articles = 1;</code>
     */
    public function setArticles(&$var)
    {
        GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Articles\Article::class);
        $this->articles = $var;
    }

}

