<?php
use Drupal\Core\DrupalKernel;
use Symfony\Component\HttpFoundation\Request;

// Specify relative path to the drupal root.
$autoloader = require_once __DIR__ . '/../../autoload.php';
$request = Request::createFromGlobals();
// Retrieve the site path
$site_path = DrupalKernel::findSitePath();
// Bootstrap drupal to different levels
$kernel = DrupalKernel::createFromRequest($request, $autoloader, 'prod');
$kernel->boot();
$kernel->handlePageCache($request);
$kernel->prepareLegacyRequest($request);
print_r($kernel);
